package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(group string) *Consumer {

	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "multiple-partitioned-topic",
		GroupID:     group,
		StartOffset: kafka.FirstOffset,
	}

	return &Consumer{Reader: kafka.NewReader(config)}
}

func (c Consumer) ListenMessages() {
	messages := make(chan *kafka.Message)

	go func() {
		for {
			message, err := c.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message", err)
			}

			messages <- &message
		}
	}()

	for message := range messages {
		log.Println("--------------------------------")
		log.Printf("consumer: %v message: %v partition:%v \n", c.Reader.Config().GroupID, string(message.Value), message.Partition)
		log.Println("--------------------------------")
	}
}
