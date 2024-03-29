package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	queue    = "header-queue"
	exchange = "header-1"
)

func DeclareHeaderExchange() {

	err := Channel.ExchangeDeclare(
		exchange,
		"headers",
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
		amqp.Table{
			"x-match": "all",
			"type":    "finance",
		})

	if err != nil {
		log.Println("RabbitMQ queue declare error :", err)
	}
}

func BindQueueToExchange(queueName string) {

	err := Channel.QueueBind(
		queueName,
		"",
		exchange,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue bind error :", err)
	}
}
