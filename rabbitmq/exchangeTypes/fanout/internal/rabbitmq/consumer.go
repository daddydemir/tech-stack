package rabbitmq

import (
	"fanout/config/rabbitmq"
	"log"
)

func Consume(count int, key string) {
	messages, err := rabbitmq.Channel.Consume(
		key,
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
		}
	}()

	<-forever
}
