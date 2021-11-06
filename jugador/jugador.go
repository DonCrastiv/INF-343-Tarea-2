package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"context"
	"log"

	pb "inf343-tarea-2/protoLiderJugador"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func getNumInput(s string, min int, max int) int32 {
	var input string
	log.Print(s)
	fmt.Scanln(&input)
	v, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Hubo un error al leer la consola: %v", err)
	}
	
	if v < min {
		return int32(min)
	} else if v > max {
		return int32(max)
	}
	return int32(v)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewJugadorClient(conn)

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
			jugada = getNumInput("Luz Roja, Luz Verde\nIngrese un número entre el 1 y el 10: ", 1, 10)
		case 2:
			jugada = getNumInput("Tirar la Cuerda\nIngrese un número entre el 1 y el 4: ", 1, 4)
		case 3:
			jugada = getNumInput("Todo o Nada\nIngrese un número entre el 1 y el 10: ", 1, 10)
		case 4:
			os.Exit(0)
		}
		rJ, err = c.EnviarJugada(ctx, &pb.JugadaToLider{Jugada: jugada, Etapa: etapa})
		if err != nil {
			log.Fatalf("Hubo un error con el envío o proceso de la jugada: %v", err)
		}
		elim = rJ.GetEliminado()
		etapa = rJ.GetEtapa()
		log.Printf("Eliminado: %t, Etapa: %d", elim, etapa)
	}
}
