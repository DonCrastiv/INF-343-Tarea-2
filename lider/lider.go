package main

import (
	"context"
	"log"
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
var etapaActual int32 = 1
var eliminados []bool
var jugoLaRonda []bool
var pasaDeEtapa []bool
var jugadaLider int32
var ipToId = make(map[net.Addr]int32)

func RemoveIndex(s []int32, index int) []int32 {
	return append(s[:index], s[index+1:]...)
}

func Etapa1 (ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error){
	return nil, nil
}

func Etapa2 (ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error){
	return nil, nil
}

func Etapa3 (ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error){
	return nil, nil
}

func (s *server) SolicitarUnirse(ctx context.Context, in *pbJugador.Unirse) (*pbJugador.RespuestaUnirse, error) {
	p, _ := peer.FromContext(ctx)
	if val, ok := ipToId[p.Addr]; !ok {
		log.Fatalf("%s ya tiene asignada la id %d", p.Addr.String(), val)
	}
	
	jugadorId++
	ipToId[p.Addr] = jugadorId
	//jugadores = append(jugadores, jugadorId) BORRAR SI NO USAMOS JUGADORES
	if jugadorId == 16 {
		log.Println("Comienza la ETAPA 1")
	}

	return &pbJugador.RespuestaUnirse{Etapa: 1}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error) {
	// Si esta se trata de la primera jugada de la ronda,
	// entonces el líder eligirá un número.
	v := false
	for i, value := range jugoLaRonda {
		v = v || (value && !eliminados[i])
	}
	if !v {
		jugadaLider = rand.Int31n(5) + 6
		// <ronda>++
	}
	// Aquí marcamos que el jugador ha enviado su jugada.
	p, _ := peer.FromContext(ctx)
	jugoLaRonda[ipToId[p.Addr] - 1] = true

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	c := pbName.NewLiderNameServiceClient(conn)
	r, err := c.EnviarJugadas(context.Background(), &pbName.Jugada{IdJugador: ipToId[p.Addr], Jugada: in.Jugada, Etapa: 1})
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
		pasaDeEtapa[ipToId[p.Addr] - 1] = true
	} else if cantidad == 4 {
		eliminado = true
	}
	if (eliminado) {
		eliminados[ipToId[p.Addr] - 1] = true
		log.Printf("Jugador %d Eliminado", ipToId[p.Addr])
	}
	
	// Se esperará hasta que todos los jugadores hayan enviado
	// su jugada para poder avanzar a la siguiente ronda.
	revisar := jugoLaRonda
	if pasaDeEtapa[ipToId[p.Addr] - 1] {
		revisar = pasaDeEtapa
	}
	for v := false; !v; {
		v = true
		for i, value := range revisar {
			v = v && (value || eliminados[i])
		}
	}
	if pasaDeEtapa[ipToId[p.Addr] - 1] {
		pasaDeEtapa[ipToId[p.Addr] - 1] = false
	}
	jugoLaRonda[ipToId[p.Addr] - 1] = false
	
	return &pbJugador.RespuestaJugada{Eliminado: eliminado, Etapa: etapa}, nil
}

func main() {
	for cont := 0; cont < 16; cont++ {
		jugoLaRonda = append(jugoLaRonda, false)
		pasaDeEtapa = append(pasaDeEtapa, false)
		eliminados = append(eliminados, false)
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
