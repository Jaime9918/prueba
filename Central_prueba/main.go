package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	prob_termino := rand.Intn(10)
	fmt.Println(prob_termino)
	if prob_termino >= 9 { //debe ser <8
		return &pb.Message{Body: "Termino"}, nil
	} else {
		return &pb.Message{Body: "aun no"}, nil
	}
}

func main() {
	qName := "Emergencias"                                           //Nombre de la cola
	hostQ := "localhost"                                             //Host de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	defer connQ.Close()
	ch, err := connQ.Channel()
	defer ch.Close()
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ

	var equipos_disp int = 2
	fmt.Println(q)
	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}
	serv := grpc.NewServer()
	for delivery := range chDelivery {
		equipos_disp = equipos_disp - 1
		fmt.Println("Pedido de ayuda de " + string(delivery.Body)) //obtiene el primer mensaje de la cola
		var validador int = 0
		if equipos_disp >= 0 {
			if validador == 0 {
				validador = 1
				fmt.Println("Envio ayuda a" + string(delivery.Body))
				//break //borrar
			}
			time.Sleep(5 * time.Second) //espera de 5 segundos
		}
		if validador == 2 {
			break
		}
		//for true {
		//	go func() {

		fmt.Println("como funciona esto?")
		pb.RegisterMessageServiceServer(serv, &server{})
		//go func() {
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
		//	}()
		//}
		fmt.Println("Esperando Emergencias")
		if validador == 1 {
			break
		}
	}
	fmt.Println("se logra x2")
}
