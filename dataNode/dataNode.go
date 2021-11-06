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
	port = ":50053"
)

type DataNodeServer struct {
	pb.UnimplementedNameDataServiceServer
}

func (s *DataNodeServer) RegistrarJugadas(ctx context.Context, in *pb.JugadaToData) (*pb.RespuestaJugada, error) {
	log.Printf("Input - num: %d | st: %d\n", in.IdJugador, in.Etapa)
	var pl = JugadaDN{idJugador: in.IdJugador, etapa: in.Etapa}
	var jgs []int32
	guardarJugada(pl.idJugador, pl.jugada, pl.etapa)
	jgs = obtenerJugada(pl.idJugador, pl.etapa)
	return &pb.RespuestaJugada{Jugadas: jgs, Cantidad: int32(len(jgs))}, nil
}

type JugadaDN struct {
	idJugador int32
	jugada    int32
	etapa     int32
}

func guardarJugada(idJugador int32, jugada int32, etapa int32) {

	filename := fmt.Sprintf("jugador_%d__etapa_%d.txt", idJugador, etapa)
	str := fmt.Sprintf("%d\n", jugada)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)
	check(err)

	f.WriteString(str)

	f.Close()
}

func obtenerJugada(idJugador int32, etapa int32) []int32 {
	filename := fmt.Sprintf("jugador_%d__etapa_%d.txt", idJugador, etapa)

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
