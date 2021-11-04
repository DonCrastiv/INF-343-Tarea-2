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

type Jugador struct {
	Id int
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewJugadorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.IngresarSolicitud(ctx, &pb.Solicitud{})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("%d", r.GetIdJugador())
}