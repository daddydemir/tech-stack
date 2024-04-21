package main

import (
	"log"
	"time"
)

func main() {

	log.Println("program is start.")

	//	createTopic()
	producer := NewProducer()
	consumer := NewConsumer()

	go consumer.listenMessages()
	producer.sendMessage("Hello world!")

	time.Sleep(time.Second * 5)
	log.Println("Program is stopped...")
}
