package main

import (
	"math/rand"
	"os"
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
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewJugadorClient(conn)

	/*
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	*/
	ctx := context.Background()
	rS, err := c.SolicitarUnirse(ctx, &pb.Unirse{})
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud: %v", err)
	}
	log.Println("Unido exitosamente al Squid Game")

	etapa := rS.GetEtapa()
	elim := false
	var rJ *pb.RespuestaJugada
	var jugada int32
	for !elim {
		switch etapa {
		case 1:
			jugada = rand.Int31n(10) + 1
		case 2:
			jugada = rand.Int31n(4) + 1
		case 3:
			jugada = rand.Int31n(10) + 1
		case 4:
			os.Exit(0)
		}
		rJ, err = c.EnviarJugada(ctx, &pb.Jugada{Jugada: jugada, Etapa: etapa})
		if err != nil {
			log.Fatalf("Hubo un error con el envío o proceso de la jugada: %v", err)
		}
		elim = rJ.GetEliminado()
		etapa = rJ.GetEtapa()
		log.Printf("Eliminado: %t, Etapa: %d", elim, etapa)
	}
}
