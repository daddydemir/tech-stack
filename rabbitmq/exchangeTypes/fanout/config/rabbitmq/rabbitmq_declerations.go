package rabbitmq

import "log"

var (
	queue    = "public-queue"
	exchange = "public-1"
)

func DeclareFanoutExchange() {

	err := Channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		true,
		false,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ exchange declare error :", err)
	}
}

func DeclareQueue() {

	_, err := Channel.QueueDeclare(
		queue,
		true,
		true,
		true,
		true,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue declare error :", err)
	}
}

func BindQueueToExchange() {

	err := Channel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue bind error :", err)
	}
}
