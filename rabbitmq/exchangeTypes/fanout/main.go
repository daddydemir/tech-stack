package main

import (
	"fanout/config/rabbitmq"
	"fanout/handler"
)

func main() {

	rabbitmq.ConnectToRabbitMQ()
	rabbitmq.DeclareFanoutExchange()

	rabbitmq.DeclareQueue("public-queue-1")
	rabbitmq.DeclareQueue("public-queue-2")
	rabbitmq.DeclareQueue("public-queue-3")

	rabbitmq.BindQueueToExchange("public-queue-1")
	rabbitmq.BindQueueToExchange("public-queue-2")
	rabbitmq.BindQueueToExchange("public-queue-3")

	handler.StartHttpServer()
}
