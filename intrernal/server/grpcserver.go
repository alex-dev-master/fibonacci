package server

import (
	"context"
	pb "github.com/alex-dev-master/fibonacci.git/proto"
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
}

func (s *BootstrapGrpc) RunRpc() error {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		return err
	}
	s.srv = grpc.NewServer()
	pb.RegisterFibonacciServer(s.srv, &fibonacciRpcServer{})
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
	res := []uint64{1, 2}
	err := validateInputFibonacci(point)
	if err != nil {
		return nil, err
	}

	return &pb.FibonacciSliceResponse{Res: res}, nil
}

func validateInputFibonacci(input *pb.FibonacciSliceRequest) error {
	if input.Y <= input.X {
		return status.Error(codes.InvalidArgument, "Y should have more than X")
	}

	return nil
}
