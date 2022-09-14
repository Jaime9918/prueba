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
	if msg.Body == "Equipo listo" {
		return &pb.Message{Body: " A PELEAAR"}, nil
	} else {
		return &pb.Message{Body: "PROBLEMAS"}, nil
	}
}

func main() {
	LabName := "Laboratiorio Santiago" //nombre del laboratorio
	qName := "Emergencias"             //nombre de la cola
	hostQ := "localhost"               //ip del servidor de RabbitMQ 172.17.0.1
	//queue_terminos := "Final guerras"                                //nombre de cola que lee terminos de guerras
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	var prob_ataque int

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
	for true {
		prob_ataque = rand.Intn(10)
		fmt.Println(prob_ataque)
		if prob_ataque > 8 { //debe ser <8
			err = ch.Publish("", qName, false, false,
				amqp.Publishing{
					Headers:     nil,
					ContentType: "text/plain",
					Body:        []byte(LabName), //Contenido del mensaje
				})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(LabName)

			listener, err := net.Listen("tcp", ":50051") //conexion sincrona
			if err != nil {
				panic("La conexion no se pudo crear" + err.Error())
			}
			serv := grpc.NewServer()
			for {
				pb.RegisterMessageServiceServer(serv, &server{})
				if err = serv.Serve(listener); err != nil {
					panic("El server no se pudo iniciar" + err.Error())
				}
			}
		} else {
			fmt.Println("Termino la guerra")
		}
		time.Sleep(5 * time.Second) //espera de 5 segundos
	}
}
