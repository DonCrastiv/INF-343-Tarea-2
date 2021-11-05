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
var solicitudes map[int32]bool
var jugadaLider int32
var ipToId = make(map[net.Addr]int32)


func (s *server) SolicitarUnirse(ctx context.Context, in *pbJugador.Unirse) (*pbJugador.RespuestaUnirse, error) {
	p, _ := peer.FromContext(ctx)
	if val, ok := ipToId[p.Addr]; !ok {
		log.Fatalf("%s ya tiene asignada la id %d", p.Addr.String(), val)
	}
	
	jugadorId++
	ipToId[p.Addr] = jugadorId
	jugadores = append(jugadores, jugadorId)
	if jugadorId == 16 {
		log.Println("Comienza la ETAPA 1")
	}

	return &pbJugador.RespuestaUnirse{Etapa: 1}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error) {
	// Si esta se trata de la primera jugada de la ronda,
	// entonces el líder eligirá un número.
	v := false
	for _, value := range solicitudes {
		v = v || value
	}
	if !v {
		jugadaLider = rand.Int31n(5) + 6
		// <ronda>++
	}
	// Aquí marcamos que el jugador ha enviado su jugada.
	p, _ := peer.FromContext(ctx)
	solicitudes[ipToId[p.Addr]] = true
	
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	c := pbName.NewLiderNameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EnviarJugadas(ctx, &pbName.Jugada{IdJugador: ipToId[p.Addr], Jugada: in.Jugada, Etapa: 1})
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud entre Lider-NameNode: %v", err)
	}
	jugadas := r.Jugadas
	// <cantidad> es el número de ronda?
	cantidad := r.Cantidad
	var suma int32 = 0

	// Reglas de eliminación del juego.
	for _, v := range jugadas {  
		suma += v  
	}
	eliminado := jugadas[cantidad-1] >= jugadaLider
	
	var etapa int32 = 1
	if suma >= 21 {
		etapa = 2
	} else if cantidad == 4 {
		eliminado = true
	}
	if (eliminado) {
		log.Printf("Jugador %d Eliminado", ipToId[p.Addr])
	}
	
	// Se esperará hasta que todos los jugadores hayan enviado
	// su jugada para poder avanzar a la siguiente ronda.
	for v := false; !v; {
		for _, value := range solicitudes {
			v = v && value
		}
	}
	solicitudes[ipToId[p.Addr]] = false
	return &pbJugador.RespuestaJugada{Eliminado: eliminado, Etapa: etapa}, nil
}

func main() {
	for cont := 0; cont < 16; cont++ {
		solicitudes[int32(cont)] = false
	}
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
