package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	Writer *kafka.Writer
}

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer() *Consumer {

	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "topic-name",
		GroupID:     "group-1",
		StartOffset: kafka.FirstOffset,
	}

	return &Consumer{kafka.NewReader(config)}
}

func NewProducer() *Producer {

	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "topic-name",
	}

	return &Producer{writer}
}

func createTopic() {

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	configs := []kafka.TopicConfig{
		{
			Topic:             "topic-name",
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(configs...)
	if err != nil {
		log.Println("create topics :", err)
	}

	log.Println("topic has been created...")
}

func (p *Producer) sendMessage(message interface{}) {
	err := p.Writer.WriteMessages(context.Background(), *prepareMessage(message))
	if err != nil {
		log.Println("sendMessage: ", err)
	} else {
		log.Println("message is send.")
	}
}

func prepareMessage(t interface{}) *kafka.Message {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Println("prepareMessage: ", err)
	}

	return &kafka.Message{
		Value: bytes,
	}
}

func (c *Consumer) listenMessages() {

	messages := make(chan *kafka.Message)

	go func() {
		for {
			message, err := c.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message: ", err)
			}
			
			messages <- &message
		}
	}()

	for message := range messages {

		log.Printf("Message received :%s \n", message.Value)

	}
}
