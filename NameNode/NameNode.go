package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pbNameData "inf343-tarea-2/protoNameData"

	"google.golang.org/grpc"
)

const (
	adress = "localhost:50051"
)

type nameNodeData struct {
	playerNumber int
	playerStage  int
	ip           string
}

type Jugada{
	
}

func (p *nameNodeData) savePlayerData() {
	filename := "nameNodeDataLocation.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)

	check(err)
	fmt.Fprintf(f, "Jugador_%d Ronda_%d 10.0.1.10\n", p.playerNumber, p.playerStage)
	f.Close()
}

func (p *nameNodeData) getStoredIP() {
	filename := "nameNodeDataLocation.txt"

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var str, ip string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		var pl, st int
		fmt.Sscanf(str, "Jugador_%d Ronda_%d %s", &pl, &st, &ip)

		if pl == p.playerNumber && st == p.playerStage {
			p.ip = ip
			break
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := grpc.Dial(adress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect\n")
	}
	defer conn.Close()
	c := pbNameData.NewNameDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RegistrarJugadas(ctx, &pbNameData.Jugada{IdJugador: int32(1), Etapa: int32(2), Jugada: int32(3)})
	//log.Printf("Jugada %v", r)
	log.Printf("%v", r.Jugadas)
}
