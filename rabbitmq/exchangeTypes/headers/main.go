package main

import (
	"headers/config/rabbitmq"
	"headers/handler"
)

func main() {

	rabbitmq.ConnectToRabbitMQ()
	rabbitmq.DeclareHeaderExchange()
	rabbitmq.DeclareQueue("header-queue")
	rabbitmq.BindQueueToExchange("header-queue")

	handler.StartHttpServer()
}
