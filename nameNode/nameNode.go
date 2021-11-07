package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	pbLiderName "inf343-tarea-2/protoLiderName"
	pbNameData "inf343-tarea-2/protoNameData"

	"google.golang.org/grpc"
)

var adress = [...]string{"localhost:50061", "localhost:50062", "localhost:50063"}
var count = 0

const (
	port = ":50052"
)

type server struct {
	pbLiderName.UnimplementedLiderNameServiceServer
}

func (s *server) EnviarJugadas(ctx context.Context, in *pbLiderName.JugadaToName) (*pbLiderName.RespuestaJugadas, error) {
	log.Printf("Recibido: %d %d %d", in.IdJugador, in.Etapa, in.Jugada)
	str := getStoredIP(in.IdJugador, in.Etapa)
	var dirDN string
	if str == "" {
		dirDN = adress[rand.Intn(2)]
		savePlayerData(in.IdJugador, in.Etapa, dirDN)
	} else {
		dirDN = str
	}
	conn, err := grpc.Dial(dirDN, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect\n")
	}
	defer conn.Close()
	c := pbNameData.NewNameDataServiceClient(conn)
	r, err := c.RegistrarJugadas(context.Background(), &pbNameData.JugadaToData{IdJugador: in.IdJugador, Etapa: in.Etapa, Jugada: in.Jugada})
	return &pbLiderName.RespuestaJugadas{Jugadas: r.Jugadas, Cantidad: r.Cantidad}, nil
}

func savePlayerData(IdJugador int32, Etapa int32, addr string) {
	filename := "nameNode/DataLocation.txt"
	if count == 0 {
		f, err := os.Create(filename)
		check(err)
		f.Close()
		count = 1
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)
	check(err)
	fmt.Fprintf(f, "Jugador_%d Ronda_%d %s\n", IdJugador, Etapa, addr)
	f.Close()
}

func getStoredIP(IdJugador int32, Etapa int32) string {
	filename := "nameNode/DataLocation.txt"

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var str, ip string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		var pl, st int
		fmt.Sscanf(str, "Jugador_%d Ronda_%d %s", &pl, &st, &ip)

		if pl == int(IdJugador) && st == int(Etapa) {
			return ip
		}
	}
	return ""
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
