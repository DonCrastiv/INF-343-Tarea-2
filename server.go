package main

import (
	"context"
	"log"
	"net"

	pb "inf343-tarea-2/protoLiderJugador"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedJugadorServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) IngresarSolicitud(ctx context.Context, in *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	log.Printf("Received: %v", in.GetParticipa())
	return &pb.RespuestaSolicitud{Message: "Solicitud recibida: " + in.GetParticipa()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
