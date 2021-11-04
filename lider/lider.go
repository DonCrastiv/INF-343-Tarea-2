package main

import (
	"context"
	"log"
	"net"
	"fmt"

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
var jugadores []int32

func (s *server) IngresarSolicitud(ctx context.Context, in *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	jugadorId++
	jugadores = append(jugadores, int32(jugadorId))
	fmt.Printf("%v", jugadores)
	return &pb.RespuestaSolicitud{Etapa: int32(1)}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pb.Jugada) (*pb.RespuestaJugada, error){
<<<<<<< HEAD


	return nil, nil
=======
	log.Println("pito")
	return &pb.RespuestaJugada{Eliminado: true, Etapa: 5}, nil
>>>>>>> bce453007ebfb1cfc9d2a186712441e15d8307c1
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
