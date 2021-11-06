package main

import (
	"math/rand"
	"time"
	"sync"
	"strconv"
	"fmt"

	"context"
	"log"

	pb "inf343-tarea-2/protoLiderJugador"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func Jugar(wg *sync.WaitGroup) {
	defer wg.Done()

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
	Id := rS.GetId()
	log.Printf("[%d] Unido exitosamente al Squid Game", Id)

	etapa := rS.GetEtapa()
	elim := false
	var rJ *pb.RespuestaJugada
	var jugada int32
	for !elim {
		time.Sleep(500 * time.Millisecond)
		switch etapa {
		case 1:
			jugada = rand.Int31n(10) + 1
		case 2:
			jugada = rand.Int31n(4) + 1
		case 3:
			jugada = rand.Int31n(10) + 1
		case 4:
			log.Printf("")
		}
		rJ, err = c.EnviarJugada(ctx, &pb.JugadaToLider{Jugada: jugada, Etapa: etapa})
		if err != nil {
			log.Fatalf("[%d] Hubo un error con el envío o proceso de la jugada: %v", Id, err)
		}
		elim = rJ.GetEliminado()
		etapa = rJ.GetEtapa()
		log.Printf("[%d] Eliminado: %t, Etapa: %d", Id, elim, etapa)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	var input string
	log.Print("Cantidad de Bots a instanciar: ")
	fmt.Scanln(&input)
	v, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Hubo un error al leer la consola: %v", err)
	}
	
	if v < 1 {
		v = 1
	} else if v > 16 {
		v = 16
	}

	var wg sync.WaitGroup
	for i := 0; i < v; i++ {
		wg.Add(1)
		go Jugar(&wg)
	}
	wg.Wait()
}

