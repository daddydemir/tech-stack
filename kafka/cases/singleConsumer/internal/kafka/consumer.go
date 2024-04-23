package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer() *Consumer {
	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "single-partitioned-topic",
		GroupID:     "single-partition-group",
		StartOffset: kafka.FirstOffset,
	}

	return &Consumer{kafka.NewReader(config)}
}

func (c Consumer) ListenMessages() {

	messages := make(chan *kafka.Message)

	go func() {
		for {
			message, err := c.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message:", err)
			}

			messages <- &message
		}
	}()

	for message := range messages {
		log.Println("--------------------------------")
		log.Printf("message: %v partition:%v \n", string(message.Value), message.Partition)
		fmt.Println(message)
		log.Println("--------------------------------")
	}
}
