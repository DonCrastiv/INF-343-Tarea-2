package main

import (
	"context"
	"log"
	"time"
	"net"

	pbJugador "inf343-tarea-2/protoLiderJugador"
	pbName "inf343-tarea-2/protoLiderName"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const (
	address = "localhost:50052"
	port = ":50051"
)

type server struct {
	pbJugador.UnimplementedJugadorServer
}

var jugadorId int32 = 0
var jugadores []int32
var ipToId = make(map[net.Addr]int32)

func (s *server) IngresarSolicitud(ctx context.Context, in *pbJugador.Solicitud) (*pbJugador.RespuestaSolicitud, error) {
	p, _ := peer.FromContext(ctx)
	if val, ok := ipToId[p.Addr]; !ok {
		log.Fatalf("%s ya tiene asignada la id %d", p.Addr.String(), val)
	}
	
	jugadorId++
	ipToId[p.Addr] = jugadorId 
	jugadores = append(jugadores, jugadorId)

	return &pbJugador.RespuestaSolicitud{Etapa: 1}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error){

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	p, _ := peer.FromContext(ctx)
	direccion := p.Addr

	c := pbName.NewLiderNameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rS, err := c.EnviarJugadas(ctx, &pbName.Jugada{IdJugador: , Jugada: , Etapa: })
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
	}

	
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grcpServer := grpc.NewServer()

	pbJugador.RegisterJugadorServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
