package main

import (  
    "fmt"
	"math/rand"
	"time"
	"context"
	"log"

	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

type Jugador struct {
	id int
}

func (j Jugador) Unirse() {

}

func (j Jugador) VerMontoAcumulado() int {
	return 100000
}

func (j Jugador) LuzRojaLuzVerde() int {
	return rand.Intn(10) + 1
}

func (j Jugador) TirarLaCuerda() int {
	return rand.Intn(4) + 1
}

func (j Jugador) TodoONada() int {
	return rand.Intn(10) + 1
}

const (
	direccionLider = "localhost:50051"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	j := Jugador{id:1}
	fmt.Println(j.LuzRojaLuzVerde())

	conn, err := grpc.Dial(direccionLider, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJugadorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 30

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
