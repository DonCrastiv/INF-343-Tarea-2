package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"net"

	pbNameData "inf343-tarea-2/protoNameData"
	pbLiderName "inf343-tarea-2/protoLiderName"

	"google.golang.org/grpc"
)

const (
	adress = "localhost:50053"
	port = ":50052"
)

type server struct {
	pbLiderName.UnimplementedLiderNameServiceServer
}

type nameNodeData struct {
	playerNumber int
	playerStage  int
	ip           string
}

func (p *nameNodeData) savePlayerData() {
	filename := "nameNodeDataLocation.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)

	check(err)
	fmt.Fprintf(f, "Jugador_%d Ronda_%d 10.0.1.10\n", p.playerNumber, p.playerStage)
	f.Close()
}

func (s *server) EnviarJugadas(ctx context.Context, in *pbLiderName.JugadaToName) (*pbLiderName.RespuestaJugadas, error) {
	log.Printf("Recibido: %d %d %d", in.IdJugador, in.Etapa, in.Jugada)
	
	conn, err := grpc.Dial(adress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect\n")
	}
	defer conn.Close()
	c := pbNameData.NewNameDataServiceClient(conn)
	r, err := c.RegistrarJugadas(context.Background(), &pbNameData.JugadaToData{IdJugador: in.IdJugador, Etapa: in.Etapa, Jugada: in.Jugada})
	return &pbLiderName.RespuestaJugadas{Jugadas: r.Jugadas, Cantidad: r.Cantidad}, nil
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

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grcpServer := grpc.NewServer()

	pbLiderName.RegisterLiderNameServiceServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
