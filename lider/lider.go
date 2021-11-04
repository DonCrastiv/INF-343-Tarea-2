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

type server struct {
	pb.UnimplementedJugadorServer
}

var jugadorId int32 = 0
func (s *server) IngresarSolicitud(ctx context.Context, in *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	log.Printf("Received: %v", in.GetParticipa())
	jugadorId++
	return &pb.RespuestaSolicitud{IdJugador: jugadorId}, nil
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
