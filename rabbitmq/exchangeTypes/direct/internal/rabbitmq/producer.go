package rabbitmq

import (
	"context"
	r "direct/config/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func PublishWithExchange(exchange, key, message string) {

	err := r.Channel.PublishWithContext(context.Background(),
		exchange,
		key,
		false,
		false,
		prepareMessage(message))

	if err != nil {
		log.Println("Error publishing : ", err)
	} else {
		log.Println("message published")
	}

}

func PublishWithQueue(queueName, message string) {

	err := r.Channel.PublishWithContext(context.Background(),
		"",
		queueName,
		false,
		false,
		prepareMessage(message))
	if err != nil {
		log.Println("Error publishing : ", err)
	}
}

func prepareMessage(message string) amqp.Publishing {

	m := amqp.Publishing{
		ContentType: "text/plain; charset=utf-8",
		Body:        []byte(message)}

	return m
}
