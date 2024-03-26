package main

import (
	"fanout/config/rabbitmq"
	"fanout/handler"
)

func main() {

	rabbitmq.ConnectToRabbitMQ()
	rabbitmq.DeclareFanoutExchange()
	rabbitmq.DeclareQueue()
	rabbitmq.BindQueueToExchange()

	handler.StartHttpServer()
}
