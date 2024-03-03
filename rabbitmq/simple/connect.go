package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Channel *amqp.Channel

func connect2rabbitmq() {
	s := "amqp://rabbit-user:password@localhost:5672"
	c, err := amqp.Dial(s)
	if err != nil {
		log.Println("RabbitMQ connection error : ", err)
	}

	Channel, err = c.Channel()
	if err != nil {
		log.Println("RabbitMQ error: ", err)
	}
}
