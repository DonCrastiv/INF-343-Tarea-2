// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoLiderJugadores

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

// JugadorClient is the client API for Jugador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JugadorClient interface {
	IngresarSolicitud(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaSolicitud, error)
	InformarEtapas(ctx context.Context, in *Etapa, opts ...grpc.CallOption) (*RespuestaEtapa, error)
	EnviarJugadas(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*RespuestaJugada, error)
	EnviarEstado(ctx context.Context, in *Estado, opts ...grpc.CallOption) (*RespuestaEstado, error)
}

type jugadorClient struct {
	cc grpc.ClientConnInterface
}

func NewJugadorClient(cc grpc.ClientConnInterface) JugadorClient {
	return &jugadorClient{cc}
}

func (c *jugadorClient) IngresarSolicitud(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaSolicitud, error) {
	out := new(RespuestaSolicitud)
	err := c.cc.Invoke(ctx, "/grpc.Jugador/IngresarSolicitud", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jugadorClient) InformarEtapas(ctx context.Context, in *Etapa, opts ...grpc.CallOption) (*RespuestaEtapa, error) {
	out := new(RespuestaEtapa)
	err := c.cc.Invoke(ctx, "/grpc.Jugador/InformarEtapas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jugadorClient) EnviarJugadas(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*RespuestaJugada, error) {
	out := new(RespuestaJugada)
	err := c.cc.Invoke(ctx, "/grpc.Jugador/EnviarJugadas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jugadorClient) EnviarEstado(ctx context.Context, in *Estado, opts ...grpc.CallOption) (*RespuestaEstado, error) {
	out := new(RespuestaEstado)
	err := c.cc.Invoke(ctx, "/grpc.Jugador/EnviarEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JugadorServer is the server API for Jugador service.
// All implementations must embed UnimplementedJugadorServer
// for forward compatibility
type JugadorServer interface {
	IngresarSolicitud(context.Context, *Solicitud) (*RespuestaSolicitud, error)
	InformarEtapas(context.Context, *Etapa) (*RespuestaEtapa, error)
	EnviarJugadas(context.Context, *Jugada) (*RespuestaJugada, error)
	EnviarEstado(context.Context, *Estado) (*RespuestaEstado, error)
	mustEmbedUnimplementedJugadorServer()
}

// UnimplementedJugadorServer must be embedded to have forward compatible implementations.
type UnimplementedJugadorServer struct {
}

func (UnimplementedJugadorServer) IngresarSolicitud(context.Context, *Solicitud) (*RespuestaSolicitud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngresarSolicitud not implemented")
}
func (UnimplementedJugadorServer) InformarEtapas(context.Context, *Etapa) (*RespuestaEtapa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InformarEtapas not implemented")
}
func (UnimplementedJugadorServer) EnviarJugadas(context.Context, *Jugada) (*RespuestaJugada, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarJugadas not implemented")
}
func (UnimplementedJugadorServer) EnviarEstado(context.Context, *Estado) (*RespuestaEstado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEstado not implemented")
}
func (UnimplementedJugadorServer) mustEmbedUnimplementedJugadorServer() {}

// UnsafeJugadorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JugadorServer will
// result in compilation errors.
type UnsafeJugadorServer interface {
	mustEmbedUnimplementedJugadorServer()
}

func RegisterJugadorServer(s grpc.ServiceRegistrar, srv JugadorServer) {
	s.RegisterService(&Jugador_ServiceDesc, srv)
}

func _Jugador_IngresarSolicitud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solicitud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JugadorServer).IngresarSolicitud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Jugador/IngresarSolicitud",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JugadorServer).IngresarSolicitud(ctx, req.(*Solicitud))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jugador_InformarEtapas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Etapa)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JugadorServer).InformarEtapas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Jugador/InformarEtapas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JugadorServer).InformarEtapas(ctx, req.(*Etapa))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jugador_EnviarJugadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JugadorServer).EnviarJugadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Jugador/EnviarJugadas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JugadorServer).EnviarJugadas(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jugador_EnviarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Estado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JugadorServer).EnviarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Jugador/EnviarEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JugadorServer).EnviarEstado(ctx, req.(*Estado))
	}
	return interceptor(ctx, in, info, handler)
}

// Jugador_ServiceDesc is the grpc.ServiceDesc for Jugador service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jugador_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Jugador",
	HandlerType: (*JugadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngresarSolicitud",
			Handler:    _Jugador_IngresarSolicitud_Handler,
		},
		{
			MethodName: "InformarEtapas",
			Handler:    _Jugador_InformarEtapas_Handler,
		},
		{
			MethodName: "EnviarJugadas",
			Handler:    _Jugador_EnviarJugadas_Handler,
		},
		{
			MethodName: "EnviarEstado",
			Handler:    _Jugador_EnviarEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liderJugador.proto",
}
