package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func createTopic() {
	_, err := Channel.QueueDeclare("queueName",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Println("error : ", err)
	}
}

func publishMessage() {
	message := amqp.Publishing{ContentType: "text/plain",
		Body: []byte("Hello, world!")}

	err := Channel.PublishWithContext(context.Background(),
		"",
		"queueName",
		false,
		false,
		message)
	if err != nil {
		log.Println("error : ", err)
	}
}

func consumeMessage() {

	messages, err := Channel.Consume("queueName",
		"",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Println("error : ", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Println("received message:  ", string(message.Body))
		}
	}()

	<-forever

}

func main() {
	connect2rabbitmq()

	consumeMessage()
	publishMessage()
}
