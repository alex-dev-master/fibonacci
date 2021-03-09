package server

import (
	"context"
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/intrernal/service"
	pb "github.com/alex-dev-master/fibonacci.git/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
)

type BootstrapGrpc struct {
	srv *grpc.Server
}

type fibonacciRpcServer struct {
	pb.UnimplementedFibonacciServer
	services *service.Service
}

func (s *BootstrapGrpc) RunRpc(services *service.Service) error {
	listener, err := net.Listen("tcp", viper.GetString("grpc.port"))
	if err != nil {
		return err
	}
	s.srv = grpc.NewServer()
	pb.RegisterFibonacciServer(s.srv, &fibonacciRpcServer{services: services})
	reflection.Register(s.srv)

	if e := s.srv.Serve(listener); e != nil {
		return e
	}

	return nil
}

func (s *BootstrapGrpc) Shutdown() {
	s.srv.GracefulStop()
}

func (receiver *fibonacciRpcServer) FibonacciSlice(ctx context.Context, point *pb.FibonacciSliceRequest) (*pb.FibonacciSliceResponse, error) {
	res, err := receiver.services.Fibonacci.GetSlice(model.Fibonacci{
		X: point.X,
		Y: point.Y,
	})
	if err != nil {
		return nil, err
	}
	err = validateInputFibonacci(point)
	if err != nil {
		return nil, err
	}

	return &pb.FibonacciSliceResponse{Res: res}, nil
}

func validateInputFibonacci(input *pb.FibonacciSliceRequest) error {
	if input.Y <= input.X {
		return status.Error(codes.InvalidArgument, "Y should have more than X")
	}	else if input.X > 92 || input.Y > 92 {
		return status.Error(codes.InvalidArgument, "Y and X must be less than 92")
	}

	return nil
}
