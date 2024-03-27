package rabbitmq

import (
	"context"
	"fanout/config/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func PublishWithQueue(queueName, message string) {

	err := rabbitmq.Channel.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		prepareMessage(message))

	if err != nil {
		log.Println("Error publishing message : ", err)
	}
}

func PublishWithExchange(exchange, message string) {

	err := rabbitmq.Channel.PublishWithContext(
		context.Background(),
		exchange,
		"",
		false,
		false,
		prepareMessage(message))
	if err != nil {
		log.Println("Error publishing message : ", err)
	}
}

func prepareMessage(message string) amqp.Publishing {

	m := amqp.Publishing{
		ContentType: "text/plain; charset=utf-8",
		Body:        []byte(message),
	}

	return m
}
