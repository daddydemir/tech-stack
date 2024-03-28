package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"headers/config/rabbitmq"
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
	} else {
		log.Printf("message send to exchange {%v} \n", exchange)
	}
}

func prepareMessage(message string) amqp.Publishing {

	m := amqp.Publishing{
		ContentType: "text/plain; charset=utf-8",
		Body:        []byte(message),
		Headers:     prepareHeader(),
	}

	return m
}

func prepareHeader() amqp.Table {
	header := amqp.Table{
		"x-match": "all",
		"type":    "message",
		"id":      1,
	}

	return header
}
