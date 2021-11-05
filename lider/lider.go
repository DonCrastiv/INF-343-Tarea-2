package main

import (
	"context"
	"log"
	"time"
	"math/rand"
	"net"

	pbJugador "inf343-tarea-2/protoLiderJugador"
	pbName "inf343-tarea-2/protoLiderName"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const (
	address = "localhost:50052"
	port = ":50051"
)

type server struct {
	pbJugador.UnimplementedJugadorServer
}

var jugadorId int32 = 0
var jugadores []int32
var ipToId = make(map[net.Addr]int32)

func (s *server) SolicitarUnirse(ctx context.Context, in *pbJugador.Unirse) (*pbJugador.RespuestaUnirse, error) {
	p, _ := peer.FromContext(ctx)
	if val, ok := ipToId[p.Addr]; !ok {
		log.Fatalf("%s ya tiene asignada la id %d", p.Addr.String(), val)
	}
	
	jugadorId++
	ipToId[p.Addr] = jugadorId 
	jugadores = append(jugadores, jugadorId)

	return &pbJugador.RespuestaUnirse{Etapa: 1}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error){

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	p, _ := peer.FromContext(ctx)
	direccion := p.Addr

	c := pbName.NewLiderNameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EnviarJugadas(ctx, &pbName.Jugada{IdJugador: ipToId[direccion], Jugada: in.Jugada, Etapa: int32(1)})
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud entre Lider-NameNode: %v", err)
	}
	jugadas := r.Jugadas
	cantidad := r.Cantidad
	var suma int32 = 0
	miJugada := rand.Int31n(5) + 6
	for _, v := range jugadas {  
			suma += v  
		}
	eliminado := jugadas[cantidad-1] >= miJugada
	
	var etapa int32 = 1
	if suma >= 21 {
		etapa = 2
	} else if cantidad == 4 {
		eliminado = true
	}
	return &pbJugador.RespuestaJugada{Eliminado: eliminado, Etapa: etapa}, nil
	
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grcpServer := grpc.NewServer()

	pbJugador.RegisterJugadorServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
