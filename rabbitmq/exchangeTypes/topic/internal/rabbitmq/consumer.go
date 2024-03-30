package rabbitmq

import (
	"log"
	"topic/config/rabbitmq"
)

func Consume(count int, queueName string) {
	messages, err := rabbitmq.Channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Println("Failed to consume : ", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("consumer {%v} : received message : %v", count, string(message.Body))
			log.Printf("message headers: %v", message.Headers)
		}
	}()

	<-forever
}
