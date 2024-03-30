package main

import (
	"topic/config/rabbitmq"
	"topic/handler"
)

func main() {

	rabbitmq.ConnectToRabbitMQ()
	rabbitmq.DeclareHeaderExchange()
	rabbitmq.DeclareQueue("topic-queue-1")
	rabbitmq.DeclareQueue("topic-queue-2")
	rabbitmq.BindQueueToExchange("topic-queue-1", "queue.*")
	rabbitmq.BindQueueToExchange("topic-queue-2", "queue.#")

	handler.StartHttpServer()
}
