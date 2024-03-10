package main

import (
	"direct/config/rabbitmq"
	"direct/handler"
)

func main() {

	rabbitmq.ConnectToRabbitMQ()
	rabbitmq.DeclareDirectExchange()
	rabbitmq.DeclareQueue()
	rabbitmq.BindQueueToExchange()

	handler.StartHttpServer()

}
