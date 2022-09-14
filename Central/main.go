package main

import (
	"context"
	"fmt"
	"log"
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

	/*if msg.Body == "Equipo listo" {
		prob_termino := rand.Intn(10)
		fmt.Println(prob_termino)
		if prob_termino > 8 { //debe ser <8
			return &pb.Message{Body: "Termino"}, nil
		} else {
			return &pb.Message{Body: "aun no termino"}, nil
		}
	} else {
		return &pb.Message{Body: "PROBLEMAS"}, nil
	}
	*/
	return &pb.Message{Body: "WWTFFFFFFFFFFFFFFFFF"}, nil
}

func main() {
	qName := "Emergencias" //Nombre de la cola
	//queue_terminos := "Final guerras"
	hostQ := "localhost" //Host de RabbitMQ 172.17.0.1
	//hostS := "localhost"                                             //Host de un Laboratorio
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	if err != nil {
		log.Fatal(err)
	}

	var equipos_disp int = 2
	fmt.Println(q)
	//fmt.Println(queue_terminos)
	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	//chDel_guerras, err := ch.Consume(queue_terminos, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	for delivery := range chDelivery {
		//port := ":50051"
		equipos_disp = equipos_disp - 1
		fmt.Println("Pedido de ayuda de " + string(delivery.Body)) //obtiene el primer mensaje de la cola

		//AQUI VA TODA LA COMUNICACION SINCRONA
		/*
			connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}

			defer connS.Close()*/
		//serviceCliente := pb.NewMessageServiceClient(connS)
		listener, err := net.Listen("tcp", ":50051") //conexion sincrona
		if err != nil {
			panic("La conexion no se pudo crear" + err.Error())
		}
		serv := grpc.NewServer()

		//for {
		pb.RegisterMessageServiceServer(serv, &server{})
		//go func() {
		fmt.Println(serv.Serve(listener))
		fmt.Println("LKDSAJKDJ44444444444444444444444444444KALSJK")
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
		var validador int = 0
		for true {
			//serviceCliente := pb.NewMessageServiceClient(connS)
			//serviceCliente := pb.NewMessageServiceClient(connS)
			//envia el mensaje al laboratorio
			fmt.Println("clkjdkajkajdkdjaskl")
			if equipos_disp >= 0 {
				if validador == 0 {
					validador = 1
					fmt.Println("Envio ayuda a" + string(delivery.Body))
					break //borrar
				}
				/*
					res, err := serviceCliente.Intercambio(context.Background(),
						&pb.Message{
							Body: "Equipo listo",
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}

					//fmt.Println(res.Body) //respuesta del laboratorio
					if res.Body == "Termino" {
						equipos_disp = equipos_disp + 1
						fmt.Println("TERMINO LA GUERRA")
						validador = 2
						defer connS.Close()
						break
					} else {
						fmt.Println("NO TERMINA AUN LA GUERRA")
					}
				*/
				time.Sleep(5 * time.Second) //espera de 5 segundos
			}
			if validador == 2 {
				break
			}
		}
		//funca := serviceCliente.nuevo_ataque()

		/*
			//time.Sleep(5 * time.Second) //espera de 5 segundos
				if equipos_disp == 0 {
					for delivery_guerras := range chDel_guerras { //funciona de manera asincronica
						equipos_disp = equipos_disp + 1
						fmt.Println("Termino la guerra en:" + string(delivery_guerras.Body))
						break
					}
				}
		*/
		fmt.Println("Esperando Emergencias")
	}
}
