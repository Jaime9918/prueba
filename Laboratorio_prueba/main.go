package main

import (
	//"context"
	"context"
	"fmt"
	"log"
	"math/rand"

	//"net"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func main() {
	LabName := "Laboratiorio Pripyat" //nombre del laboratorio
	qName := "Emergencias"            //nombre de la cola
	hostQ := "localhost"              //ip del servidor de RabbitMQ 172.17.0.1
	hostS := "localhost"
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
	port := ":50051"
	connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
	for true {
		for true {
			prob_ataque = rand.Intn(10)
			fmt.Println(prob_ataque)
			if prob_ataque > 5 { //debe ser <8
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

				defer connS.Close()
				serviceCliente := pb.NewMessageServiceClient(connS)
				for true {
					res, err := serviceCliente.Intercambio(context.Background(),
						&pb.Message{
							Body: "Equipo listo",
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}

					fmt.Println(res.Body) //respuesta del laboratorio
					if res.Body == "Termino" {
						break
					}
					time.Sleep(5 * time.Second) //espera de 5 segundo
				}
				break
			} else {
				fmt.Println("No hay ataque")
			}
			time.Sleep(5 * time.Second) //espera de 5 segundos
		}
	}
	fmt.Println("funciona de pana")
}