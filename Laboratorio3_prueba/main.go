package main

import (
	//"context"

	"fmt"
	"log"
	"math/rand"

	//"net"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	//"google.golang.org/grpc"
)

func numeroAleatorio(valorMin int, valorMax int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return valorMin + rand.Intn(valorMax-valorMin)
}

func main() {
	LabName := "Laboratorio Pohang - Korea" //nombre del laboratorio
	qName := "Emergencias"                  //nombre de la cola
	hostQ := "localhost"                    //ip del servidor de RabbitMQ 172.17.0.1
	queue_escuadron := "escuadron lab3"
	queue_retorno := "retorno"
	//hostS := "localhost"
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	defer connQ.Close()
	ch, err := connQ.Channel()
	defer ch.Close()
	q1, err := ch.QueueDeclare(queue_escuadron, false, false, false, false, nil)                 //Se crea la cola en RabbitMQ
	chDelivery_escuadron, err := ch.Consume(queue_escuadron, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()
	fmt.Println(q1)
	defer ch.Close()
	//port := ":50051"
	//connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
	for true {
		prob_ataque := numeroAleatorio(1, 10)
		if prob_ataque < 8 { //debe ser <8
			fmt.Println("Analizando estado " + LabName + ": [ESTALLIDO]")
			err = ch.Publish("", qName, false, false,
				amqp.Publishing{
					Headers:     nil,
					ContentType: "text/plain",
					Body:        []byte(LabName), //Contenido del mensaje
				})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("SOS enviado a la central. Esperando respuesta...")
			time.Sleep(3 * time.Second) //espera de 1 segundo
			escuadron := "-1"
			for delivery := range chDelivery_escuadron {
				escuadron = string(delivery.Body)
				fmt.Println("Llega escuadron " + escuadron + " conteniendo estallido...")
				break
			}
			for true {
				prob_termino := numeroAleatorio(1, 10)
				if prob_termino <= 6 {
					fmt.Println("Revisando estado de la resolucion: [LISTO]")
					err = ch.Publish("", queue_retorno, false, false,
						amqp.Publishing{
							Headers:     nil,
							ContentType: "text/plain",
							Body:        []byte(string(escuadron)), //Contenido del mensaje
						})
					fmt.Println("Estallido contenido, Escuadrón " + escuadron + " retornando")
					time.Sleep(2 * time.Second) //espera de 1 segundo
					break
				} else {
					fmt.Println("Revisando estado de la resolucion: [NO LISTO]")
					time.Sleep(5 * time.Second) //espera de 1 segundo
				}
			}
			//defer connS.Close()
			//serviceCliente := pb.NewMessageServiceClient(connS)
			/*
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
			*/
			//break
		} else {
			fmt.Println("Analizando estado " + LabName + ": [OK]")
		}
		time.Sleep(5 * time.Second) //espera de 5 segundos
	}
	//fmt.Println("funciona de pana")
}
