package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "inf343-tarea-2/protoNameData"

	"google.golang.org/grpc"
)

const (
	port = ":50063"
)

var l_jugadores []int32

type DataNodeServer struct {
	pb.UnimplementedNameDataServiceServer
}

func (s *DataNodeServer) RegistrarJugadas(ctx context.Context, in *pb.JugadaToData) (*pb.RespuestaJugada, error) {
	log.Printf("Input - num: %d | st: %d\n", in.IdJugador, in.Etapa)
	var jgs []int32
	guardarJugada(in.IdJugador, in.Jugada, in.Etapa)
	jgs = obtenerJugada(in.IdJugador, in.Etapa)
	return &pb.RespuestaJugada{Jugadas: jgs, Cantidad: int32(len(jgs))}, nil
}

func valueInSlice(value int32, list []int32) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	return false
}

func guardarJugada(idJugador int32, jugada int32, etapa int32) {
	filename := fmt.Sprintf("dataNode/jugador_%d__etapa_%d.txt", idJugador, etapa)
	str := fmt.Sprintf("%d\n", jugada)

	if valueInSlice(idJugador, l_jugadores) {
		f, err := os.OpenFile(filename, os.O_APPEND, 0600)
		check(err)
		f.WriteString(str)
		f.Close()
	} else {
		l_jugadores = append(l_jugadores, idJugador)
		f, err := os.Create(filename)
		check(err)
		f.WriteString(str)
		f.Close()
	}

}

func obtenerJugada(idJugador int32, etapa int32) []int32 {
	filename := fmt.Sprintf("dataNode/jugador_%d__etapa_%d.txt", idJugador, etapa)

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var num int32
	var str string
	var pl []int32
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		fmt.Sscanf(str, "%d", &num)
		pl = append(pl, num)

	}
	return pl
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fatal Error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameDataServiceServer(s, &DataNodeServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal Error: %v", err)
	}
}
