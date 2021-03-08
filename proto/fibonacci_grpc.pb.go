// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FibonacciClient is the client API for Fibonacci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibonacciClient interface {
	FibonacciSlice(ctx context.Context, in *FibonacciSliceRequest, opts ...grpc.CallOption) (*FibonacciSliceResponse, error)
}

type fibonacciClient struct {
	cc grpc.ClientConnInterface
}

func NewFibonacciClient(cc grpc.ClientConnInterface) FibonacciClient {
	return &fibonacciClient{cc}
}

func (c *fibonacciClient) FibonacciSlice(ctx context.Context, in *FibonacciSliceRequest, opts ...grpc.CallOption) (*FibonacciSliceResponse, error) {
	out := new(FibonacciSliceResponse)
	err := c.cc.Invoke(ctx, "/proto.Fibonacci/FibonacciSlice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibonacciServer is the server API for Fibonacci service.
// All implementations must embed UnimplementedFibonacciServer
// for forward compatibility
type FibonacciServer interface {
	FibonacciSlice(context.Context, *FibonacciSliceRequest) (*FibonacciSliceResponse, error)
	mustEmbedUnimplementedFibonacciServer()
}

// UnimplementedFibonacciServer must be embedded to have forward compatible implementations.
type UnimplementedFibonacciServer struct {
}

func (UnimplementedFibonacciServer) FibonacciSlice(context.Context, *FibonacciSliceRequest) (*FibonacciSliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FibonacciSlice not implemented")
}
func (UnimplementedFibonacciServer) mustEmbedUnimplementedFibonacciServer() {}

// UnsafeFibonacciServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibonacciServer will
// result in compilation errors.
type UnsafeFibonacciServer interface {
	mustEmbedUnimplementedFibonacciServer()
}

func RegisterFibonacciServer(s grpc.ServiceRegistrar, srv FibonacciServer) {
	s.RegisterService(&Fibonacci_ServiceDesc, srv)
}

func _Fibonacci_FibonacciSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibonacciSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServer).FibonacciSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Fibonacci/FibonacciSlice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServer).FibonacciSlice(ctx, req.(*FibonacciSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fibonacci_ServiceDesc is the grpc.ServiceDesc for Fibonacci service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fibonacci_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Fibonacci",
	HandlerType: (*FibonacciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FibonacciSlice",
			Handler:    _Fibonacci_FibonacciSlice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fibonacci.proto",
}
