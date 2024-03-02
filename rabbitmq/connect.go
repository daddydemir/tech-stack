package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Channel *amqp.Channel

func connect2rabbitmq(){
	c , err := amqp.Dial("connection_string")
	if  err != nil {
		log.Println("RabbitMQ connection error : ", err)
	}

	Channel, err = c.Channel()
	if err != nil {
		log.Println("RabbitMQ error: ", err)
	}
}
