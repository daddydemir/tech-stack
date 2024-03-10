package rabbitmq

import "log"

var (
	queue    = "queue"
	exchange = "personal-2"
)

func DeclareDirectExchange() {

	err := Channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		true,
		false,
		true,
		nil)

	if err != nil {
		log.Println("RabbitMQ exchange declare error :", err)
	}
}

func DeclareQueue() {

	_, err := Channel.QueueDeclare(queue,
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

	err := Channel.QueueBind(queue,
		"",
		exchange,
		false,
		nil)
	if err != nil {
		log.Println("RabbitMQ queue bind error :", err)
	}
}
