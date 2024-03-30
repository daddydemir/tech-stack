package rabbitmq

import (
	"log"
)

var (
	queue    = "topic-queue"
	exchange = "topic-1"
)

func DeclareHeaderExchange() {

	err := Channel.ExchangeDeclare(
		exchange,
		"topic",
		true,
		true,
		false,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ exchange declare error :", err)
	}
}

func DeclareQueue(queueName string) {

	_, err := Channel.QueueDeclare(
		queueName,
		true,
		true,
		true,
		true,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue declare error :", err)
	}
}

func BindQueueToExchange(queueName, key string) {

	err := Channel.QueueBind(
		queueName,
		key,
		exchange,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue bind error :", err)
	}
}
