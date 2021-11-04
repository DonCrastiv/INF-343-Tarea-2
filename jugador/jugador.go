package main

import (
	//"math/rand"
	"time"

	"context"
	"log"

	pb "inf343-tarea-2/protoLiderJugador"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewJugadorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rS, err := c.IngresarSolicitud(ctx, &pb.Solicitud{})
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
	}
	etapa := rS.GetEtapa()
	elim := false
	var rJ *pb.RespuestaJugada
	var jugada int32
	for elim == false {
		switch etapa {
		case 1:
			jugada = 1
		case 2:
			jugada = 1
		case 3:
			jugada = 1
		}
		rJ, err = c.EnviarJugada(ctx, &pb.Jugada{Jugada: jugada})
		etapa = rJ.GetEtapa()
		elim = rJ.GetEliminado()
	}
}
