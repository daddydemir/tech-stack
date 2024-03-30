package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"topic/config/rabbitmq"
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
		"queue.news",
		false,
		false,
		prepareMessage(message))
	if err != nil {
		log.Println("Error publishing message : ", err)
	} else {
		log.Printf("message send to exchange {%v} \n", exchange)
	}
}

func prepareMessage(message string) amqp.Publishing {

	m := amqp.Publishing{
		ContentType: "text/plain; charset=utf-8",
		Body:        []byte(message),
	}

	return m
}
