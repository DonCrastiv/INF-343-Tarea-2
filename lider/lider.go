package main

import (
	"context"
	"log"
	"math/rand"
	"math"
	"net"
	"time"
	"sync"
	"strconv"


	pbJugador "inf343-tarea-2/protoLiderJugador"
	pbName "inf343-tarea-2/protoLiderName"
	amqp "github.com/rabbitmq/amqp091-go"

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

var cochinoCandado sync.Mutex

var eliminados [players]bool
var ipToId = make(map[net.Addr]int32)
var jugadasLider [6]int32
var iterador [players]int32

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func updatePozo(idJugador int32, etapa int32){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body1 := strconv.Itoa(int(idJugador))
	body2 := strconv.Itoa(int(etapa))
	body := body1 + " " + body2
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)

}

func LuzRojaLuzVerde(idJugador int32, jugada int32) (bool, int32) {
	log.Printf("El jugador %d ha sacado un %d", idJugador, jugada)
	
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	
	c := pbName.NewLiderNameServiceClient(conn)
	r, err := c.EnviarJugadas(context.Background(), &pbName.JugadaToName{IdJugador: idJugador, Jugada: jugada, Etapa: 1})
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
	
	// PROVISIONAL
	jugadaLider := jugadasLider[ronda - 1]
	// END PROVISIONAL

	eliminado := jugada >= jugadaLider
	log.Printf("Jugador %d: %d, Lider: %d, Suma jugador %d: %d", idJugador, jugada, jugadaLider, idJugador, suma)
	
	var etapa int32 = 1
	if suma >= 21 {	
		etapa = 2
	} else if ronda == 4 {
		eliminado = true
	} else if ronda > 4 {
		log.Fatalf("El jugador %d ha jugado más de 4 rondas (%d)", idJugador, ronda)
	}

	return eliminado, etapa
}

var jugadasTC [players]int32
var elegidoTC int32 = 0
var equipoEliminado int
func TirarCuerda(idJugador int32, jugada int32) (bool, int32) {
	log.Printf("El jugador %d ha sacado un %d", idJugador, jugada)
	jugadasTC[idJugador - 1] = jugada
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
	if vivos == 1 {
		return false, 4
	}
	
	// Eliminamos al jugador al azar si son impares
	if (vivos%2 != 0) {
		cochinoCandado.Lock()
		if elegidoTC == 0 {
			for i := rand.Intn(vivos); i >= 0; elegidoTC++ {
				if !eliminados[elegidoTC] {
					i -= 1
				}
			}
			eliminados[elegidoTC - 1] = true
			log.Printf("Ya que la cantidad de jugadores es impar, el jugador %d será eliminado", elegidoTC)
		}
		cochinoCandado.Unlock()
		vivos--
	}
	if idJugador == elegidoTC {
		return true, 2
	}

	// Dividimos a los jugadores vivos en la mitad por orden
	// de id y calculamos el valor sumado de los equipos.
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

	// Eliminamos de acuerdo a la paridad.
	pariLider := jugadasLider[4]%2
	log.Printf("Suma equipo 1: %d, Suma equipo 2: %d, Líder: %d", sum[0], sum[1], jugadasLider[4])
	if (sum[0]%2 != pariLider) && (sum[1]%2 != pariLider) {
		log.Printf("Ningún equipo tiene la misma paridad que el Líder. El equipo %d queda eliminado", equipoEliminado)
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
var rivalesTN [players]int32
var elegidoTN int32 = 0
func TodoNada(idJugador int32, jugada int32) (bool, int32) {
	log.Printf("El jugador %d ha sacado un %d", idJugador, jugada)
	jugadasTN[idJugador - 1] = jugada
	
	vivos := 0
	for i := range jugadasTN {
		if !eliminados[i]  {
			vivos++
		}
	}
	if vivos == 1 {
		return false, 4
	}

	// Esperar a que todos hayan hecho su jugada.
	for listo := false; !listo; {
		listo = true
		for i, value := range jugadasTN {
			if !eliminados[i] && (value == 0) {
				listo = false
				break
			}
		}
	}

	if (vivos%2 != 0) {
		cochinoCandado.Lock()
		if elegidoTN == 0 {
			for i := rand.Intn(vivos); i >= 0; elegidoTN++ {
				if !eliminados[elegidoTN] {
					i -= 1
				}
			}
			eliminados[elegidoTN - 1] = true
			log.Printf("Ya que la cantidad de jugadores es impar, el jugador %d será eliminado", elegidoTN)
		}
		cochinoCandado.Unlock()
		vivos--
	}
	if idJugador == elegidoTN {
		return true, 3
	}

	// Asigno las parejas
	cochinoCandado.Lock()
	if rivalesTN[idJugador - 1] == -1 {
		for rivalIndice, value := range rivalesTN {
			// Asigno como rival al primer jugador no muerto y sin pareja que encuentre
			if !eliminados[rivalIndice] && (value == -1) && (idJugador != int32(rivalIndice) + 1) {
				log.Printf("El jugador %d y el jugador %d son pareja", idJugador, rivalIndice + 1)
				rivalesTN[idJugador - 1] = int32(rivalIndice)
				rivalesTN[rivalIndice] = idJugador - 1
			}
		}
	}
	cochinoCandado.Unlock()

	log.Printf("Jugador %d (%d) vs Jugador %d (%d)", idJugador, int(math.Abs(float64(jugadasLider[5] - jugada))), rivalesTN[idJugador - 1] + 1, int(math.Abs(float64(jugadasLider[5] - jugadasTN[rivalesTN[idJugador - 1]]))))
	if math.Abs(float64(jugadasLider[5] - jugada)) <= math.Abs(float64(jugadasLider[5] - jugadasTN[rivalesTN[idJugador - 1]])) {
		return false, 4
	}
	return true, 3
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

func (s *server) EnviarJugada(ctx context.Context, in *pbJugador.JugadaToLider) (*pbJugador.RespuestaJugada, error) {
	p, _ := peer.FromContext(ctx)
	var eliminado bool
	var etapa int32
	switch in.Etapa {
	case 1:
		eliminado, etapa = LuzRojaLuzVerde(ipToId[p.Addr], in.Jugada)
	case 2:
		log.Printf("El jugador %d sobrevivió y avanza a la ETAPA 2", ipToId[p.Addr])
		eliminado, etapa = TirarCuerda(ipToId[p.Addr], in.Jugada)
	case 3:
		log.Printf("El jugador %d sobrevivió y avanza a la ETAPA 3", ipToId[p.Addr])
		eliminado, etapa = TodoNada(ipToId[p.Addr], in.Jugada)
	}
	
	if (eliminado) {
		eliminados[ipToId[p.Addr] - 1] = true
		log.Printf("Jugador %d Eliminado", ipToId[p.Addr])
	}
	if (etapa == 4) {
		log.Printf("El jugador %d ha ganado el Squid Game", ipToId[p.Addr])
	}
	return &pbJugador.RespuestaJugada{Eliminado: eliminado, Etapa: etapa}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := range iterador {
		eliminados[i] = false
		jugadasTC[i] = 0
		jugadasTN[i] = 0
		rivalesTN[i] = -1
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

