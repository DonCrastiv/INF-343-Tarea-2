// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoLiderPozo

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

// LiderPozoServiceClient is the client API for LiderPozoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderPozoServiceClient interface {
	ConsultarPozo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type liderPozoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderPozoServiceClient(cc grpc.ClientConnInterface) LiderPozoServiceClient {
	return &liderPozoServiceClient{cc}
}

func (c *liderPozoServiceClient) ConsultarPozo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.LiderPozoService/ConsultarPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderPozoServiceServer is the server API for LiderPozoService service.
// All implementations must embed UnimplementedLiderPozoServiceServer
// for forward compatibility
type LiderPozoServiceServer interface {
	ConsultarPozo(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedLiderPozoServiceServer()
}

// UnimplementedLiderPozoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiderPozoServiceServer struct {
}

func (UnimplementedLiderPozoServiceServer) ConsultarPozo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarPozo not implemented")
}
func (UnimplementedLiderPozoServiceServer) mustEmbedUnimplementedLiderPozoServiceServer() {}

// UnsafeLiderPozoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderPozoServiceServer will
// result in compilation errors.
type UnsafeLiderPozoServiceServer interface {
	mustEmbedUnimplementedLiderPozoServiceServer()
}

func RegisterLiderPozoServiceServer(s grpc.ServiceRegistrar, srv LiderPozoServiceServer) {
	s.RegisterService(&LiderPozoService_ServiceDesc, srv)
}

func _LiderPozoService_ConsultarPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderPozoServiceServer).ConsultarPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderPozoService/ConsultarPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderPozoServiceServer).ConsultarPozo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// LiderPozoService_ServiceDesc is the grpc.ServiceDesc for LiderPozoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiderPozoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.LiderPozoService",
	HandlerType: (*LiderPozoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConsultarPozo",
			Handler:    _LiderPozoService_ConsultarPozo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liderPozo.proto",
}