package main

import (
	"log"
	"context"
	"net"
	"strings"
	"os"
	"strconv"
	pb "inf343-tarea-2/protoLiderPozo"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

const (
	port = ":50054"
)

var monto int32 = 0

type server struct {
	pb.UnimplementedLiderPozoServiceServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConexionGRPC () (){

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grcpServer := grpc.NewServer()

	pb.RegisterLiderPozoServiceServer(grcpServer, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) ConsultarPozo (ctx context.Context, in *pb.Request) (*pb.Response, error) {

	return &pb.Response{Monto: monto}, nil
}



func main() {

	file, err := os.Create("pozo.txt")

	go ConexionGRPC()
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var mensaje = strings.Split(string(d.Body), " ")
			var idJugador = mensaje[0]
			var etapa = mensaje[1]
			monto = monto + int32(100000000)
			file.WriteString("Jugador_"+idJugador+" Ronda_"+etapa+" "+strconv.Itoa(int(monto)))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
