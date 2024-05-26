package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer() *Producer {

	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "multiple-partitioned-topic",
	}

	return &Producer{Writer: writer}
}

func (p Producer) SendMessage(message string) {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling: ", err)
	}

	msg := kafka.Message{
		Value: bytes,
	}

	err = p.Writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println("Error writing: ", err)
	}
}
