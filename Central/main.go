package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	//"net"
	"time"

	pb "github.com/Jaime9918/prueba/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	//"google.golang.org/grpc"
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
	qName := "Emergencias" //Nombre de la cola
	hostQ := "dist014"     //Host de RabbitMQ 172.17.0.1
	queue_retorno := "retorno"
	queue_escuadron1 := "escuadron lab1"
	queue_escuadron2 := "escuadron lab2"
	queue_escuadron3 := "escuadron lab3"
	queue_escuadron4 := "escuadron lab4"
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	defer connQ.Close()
	ch, err := connQ.Channel()
	defer ch.Close()
	//--------------------------------------
	//------------------------------------------------
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)          //Se crea la cola en RabbitMQ
	q1, err := ch.QueueDeclare(queue_retorno, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	var equipos_disp int = 2
	fmt.Println(q)
	fmt.Println(q1)
	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil)                 //obtiene la cola de RabbitMQ
	chDelivery_retorno, err := ch.Consume(queue_retorno, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	/*listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}
	serv := grpc.NewServer()
	if err = serv.Serve(listener); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}*/
	equipo_1 := 0 //equipo sin usar:0, equipo usado 1
	file, err := os.Create("SOLICITUDES.txt")
	defer file.Close()

	//equipo_2:= 0 		//equipo sin usar:0, equipo usado 1
	equipo := 0
	nombre_lab := "xd"
	num_escuadra := "2"
	cant_consultas := "1"
	for delivery := range chDelivery {
		fmt.Println("Mensaje asíncrono de " + string(delivery.Body)) //obtiene el primer mensaje de la cola
		var validador int = 0
		contador := 0
		for equipos_disp <= 0 {
			for delivery2 := range chDelivery_retorno {
				switch contador {
				case 0:
					//fmt.Println("Retorno la escuadra " + string(delivery2.Body) + " a la central")
					num_escuadra = string(delivery2.Body)
					equipos_disp = equipos_disp + 1
					if string(delivery2.Body) == "1" {
						equipo_1 = 0
					}
				case 1:
					nombre_lab = string(delivery2.Body)
				case 2:
					cant_consultas = string(delivery2.Body)
					fmt.Println("Retorno la escuadra " + num_escuadra + " a la central desde el " + nombre_lab)
				}
				if contador == 2 {
					file.WriteString(nombre_lab + ";" + cant_consultas + "\n")
					break
				}
				contador++
			}
			break
		}
		if equipo_1 == 0 {
			equipo = 1
		} else {
			equipo = 2
		}
		if equipos_disp >= 0 {
			equipos_disp = equipos_disp - 1
			if validador == 0 {
				validador = 1
				fmt.Println("Se envía Escuadron numero " + strconv.Itoa(equipo) + " a " + string(delivery.Body))
				switch string(delivery.Body) {
				case "Laboratorio Pripiat":
					err = ch.Publish("", queue_escuadron1, false, false,
						amqp.Publishing{
							Headers:     nil,
							ContentType: "text/plain",
							Body:        []byte(strconv.Itoa(equipo)), //Contenido del mensaje
						})
				case "Laboratorio Renca - Chile":
					err = ch.Publish("", queue_escuadron2, false, false,
						amqp.Publishing{
							Headers:     nil,
							ContentType: "text/plain",
							Body:        []byte(strconv.Itoa(equipo)), //Contenido del mensaje
						})
				case "Laboratorio Pohang - Korea":
					err = ch.Publish("", queue_escuadron3, false, false,
						amqp.Publishing{
							Headers:     nil,
							ContentType: "text/plain",
							Body:        []byte(strconv.Itoa(equipo)), //Contenido del mensaje
						})
				case "Laboratorio Kampala - Uganda":
					err = ch.Publish("", queue_escuadron4, false, false,
						amqp.Publishing{
							Headers:     nil,
							ContentType: "text/plain",
							Body:        []byte(strconv.Itoa(equipo)), //Contenido del mensaje
						})
				}
				equipo_1 = 1
				if err != nil {
					log.Fatal(err)
				}
			}
			time.Sleep(5 * time.Second) //espera de 5 segundos
		}
		//	if validador == 2 {
		//		break
		//	}
		//fmt.Println("como funciona esto?")
		//pb.RegisterMessageServiceServer(serv, &server{})
		//go func() {
		//if err = serv.Serve(listener); err != nil {
		//	panic("El server no se pudo iniciar" + err.Error())
		//}
		//	}()
	}
}
