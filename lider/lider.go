package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pbJugador "inf343-tarea-2/protoLiderJugador"
	//	pbName "inf343-tarea-2/protoLiderName"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const (
	address = "localhost:50052"
	port = ":50051"
	players = 2
)

type server struct {
	pbJugador.UnimplementedJugadorServer
}

var eliminados [2]bool
//var pasaDeRonda [2]bool
//var pasaDeEtapa [2]bool
var ipToId = make(map[net.Addr]int32)
var jugadasLider [6]int32

func RemoveIndex(s []int32, index int) []int32 {
	return append(s[:index], s[index+1:]...)
}

// PROVISIONAL
var suma [players]int32
var ronda [players]int32
// END PROVISIONAL
func LuzRojaLuzVerde(idJugador int32, jugada int32) (bool, int32) {
	/*
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	c := pbName.NewLiderNameServiceClient(conn)
	r, err := c.EnviarJugadas(context.Background(), &pbName.Jugada{IdJugador: idJugador, Jugada: jugada, Etapa: 1})
	if err != nil {
		log.Fatalf("Hubo un error con el envío o proceso de la solicitud entre Lider-NameNode: %v", err)
	}
	jugadas := r.Jugadas
	// <cantidad> es el número de ronda?
	ronda := r.Cantidad
	var suma int32 = 0

	// Reglas de eliminación del juego.
	for _, v := range jugadas {  
		suma += v  
	}
	*/
	// PROVISIONAL
	suma[idJugador - 1] += jugada
	suma := suma[idJugador - 1]
	ronda[idJugador - 1] += 1
	ronda := ronda[idJugador - 1]
	jugadaLider := jugadasLider[ronda - 1]
	// END PROVISIONAL

	eliminado := jugada >= jugadaLider
	log.Printf("Jugador %d: %d, Lider: %d, Suma jugador %d: %d", idJugador, jugada, jugadaLider, idJugador, suma)
	
	var etapa int32 = 1
	if suma >= 21 {	
		etapa = 2
		//pasaDeEtapa[idJugador - 1] = true
	} else if ronda == 4 {
		eliminado = true
	} else if ronda > 4 {
		log.Fatalf("El jugador %d ha jugado más de 4 rondas (%d)", idJugador, ronda)
	}

	return eliminado, etapa
}

var jugadasTC [players]int32
var elegido int32 = 0
var equipoEliminado int
func TirarCuerda(idJugador int32, jugada int32) (bool, int32) {
	jugadasTC[idJugador] = jugada

	// Esperar a que todos hayan hecho su jugada.
	// Aprovechamos para contar los jugadores vivos.
	var vivos int
	for listo := false; !listo; {
		vivos = players
		listo = true
		for i, value := range jugadasTC {
			if eliminados[i] {
				vivos--
			} else if value == 0 {
				listo = false
			}
		}
	}
	// Eliminamos al jugador al azar si son impares
	if (vivos%2 != 0) {
		if elegido == 0 {
			for i := rand.Intn(vivos); i >= 0; elegido++ {
				if !eliminados[elegido] {
					i -= 1
				}
			}
			eliminados[elegido - 1] = true
		}
		vivos--
	}
	if idJugador == elegido {
		return true, 2
	}

	equipo := vivos/2
	cont := 0
	var miEquipo int
	sum := []int32{0, 0}
	for i, value := range jugadasTC {
		if !eliminados[i] {
			cont++
			if cont <= equipo {
				sum[0] += value
				if idJugador == int32(i + 1) {
					miEquipo = 1
				}
			} else {
				sum[1] += value
				if idJugador == int32(i + 1) {
					miEquipo = 2
				}
			}
		}
	}

	pariLider := jugadasLider[4]%2
	if (sum[0]%2 != pariLider) && (sum[1]%2 != pariLider) {
		if miEquipo == equipoEliminado {
			return true, 2
		}
		return false, 3
	} else if sum[miEquipo - 1]%2 == pariLider {
		return false, 3
	}
	return true, 2
}

var jugadasTN [players]int32
func TodoNada(idJugador int32, jugada int32) (bool, int32) {
	jugadasTN[idJugador] = jugada
	return true, 0
}

var jugadorId int32 = 0
func (s *server) SolicitarUnirse(ctx context.Context, in *pbJugador.Unirse) (*pbJugador.RespuestaUnirse, error) {
	p, _ := peer.FromContext(ctx)
	if val, ok := ipToId[p.Addr]; ok {
		log.Fatalf("%s ya tiene asignada la id %d", p.Addr.String(), val)
	}
	jugadorId++
	ipToId[p.Addr] = jugadorId
	log.Printf("A la ip %s se le ha asignado la id %d", p.Addr.String(), jugadorId)

	if jugadorId == 16 {
		log.Println("Comienza la ETAPA 1")
	}

	return &pbJugador.RespuestaUnirse{Etapa: 1}, nil
}

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.Jugada) (*pbJugador.RespuestaJugada, error) {
	p, _ := peer.FromContext(ctx)
	log.Printf("El jugador %d ha sacado un %d", ipToId[p.Addr], in.Jugada)
	var eliminado bool
	var etapa int32
	switch in.Etapa {
	case 1:
		eliminado, etapa = LuzRojaLuzVerde(ipToId[p.Addr], in.Jugada)
	case 2:
		eliminado, etapa = TirarCuerda(ipToId[p.Addr], in.Jugada)
	case 3:
		eliminado, etapa = TodoNada(ipToId[p.Addr], in.Jugada)
	}
	
	if (eliminado) {
		eliminados[ipToId[p.Addr] - 1] = true
		log.Printf("Jugador %d Eliminado", ipToId[p.Addr])
	}
	
	return &pbJugador.RespuestaJugada{Eliminado: eliminado, Etapa: etapa}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := range suma {
		/*
		pasaDeRonda[i] = false
		pasaDeEtapa[i] = false
		eliminados[i] = false
		*/
		jugadasTC[i] = 0
		jugadasTN[i] = 0

		// PROVISIONAL
		suma[i] = 0
		ronda[i] = 0
		// END PROVISIONAL
	}
	for i := 0; i < 4; i++ {
		jugadasLider[i] = rand.Int31n(5) + 6
	}
	jugadasLider[4] = rand.Int31n(4) + 1
	jugadasLider[5] = rand.Int31n(10) + 1

	equipoEliminado = rand.Intn(2) + 1

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
