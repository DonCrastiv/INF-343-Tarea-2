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
	log.Printf("Received")
	return &pb.RespuestaSolicitud{IdJugador: 2}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grcpServer := grpc.NewServer()

	pb.RegisterJugadorServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
