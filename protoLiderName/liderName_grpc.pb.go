// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoLiderName

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

// LiderNameServiceClient is the client API for LiderNameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderNameServiceClient interface {
	EnviarJugadas(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*RespuestaJugadas, error)
}

type liderNameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderNameServiceClient(cc grpc.ClientConnInterface) LiderNameServiceClient {
	return &liderNameServiceClient{cc}
}

func (c *liderNameServiceClient) EnviarJugadas(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*RespuestaJugadas, error) {
	out := new(RespuestaJugadas)
	err := c.cc.Invoke(ctx, "/grpc.LiderNameService/EnviarJugadas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderNameServiceServer is the server API for LiderNameService service.
// All implementations must embed UnimplementedLiderNameServiceServer
// for forward compatibility
type LiderNameServiceServer interface {
	EnviarJugadas(context.Context, *Jugada) (*RespuestaJugadas, error)
	mustEmbedUnimplementedLiderNameServiceServer()
}

// UnimplementedLiderNameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiderNameServiceServer struct {
}

func (UnimplementedLiderNameServiceServer) EnviarJugadas(context.Context, *Jugada) (*RespuestaJugadas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarJugadas not implemented")
}
func (UnimplementedLiderNameServiceServer) mustEmbedUnimplementedLiderNameServiceServer() {}

// UnsafeLiderNameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderNameServiceServer will
// result in compilation errors.
type UnsafeLiderNameServiceServer interface {
	mustEmbedUnimplementedLiderNameServiceServer()
}

func RegisterLiderNameServiceServer(s grpc.ServiceRegistrar, srv LiderNameServiceServer) {
	s.RegisterService(&LiderNameService_ServiceDesc, srv)
}

func _LiderNameService_EnviarJugadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderNameServiceServer).EnviarJugadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderNameService/EnviarJugadas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderNameServiceServer).EnviarJugadas(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

// LiderNameService_ServiceDesc is the grpc.ServiceDesc for LiderNameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiderNameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.LiderNameService",
	HandlerType: (*LiderNameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarJugadas",
			Handler:    _LiderNameService_EnviarJugadas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liderName.proto",
}