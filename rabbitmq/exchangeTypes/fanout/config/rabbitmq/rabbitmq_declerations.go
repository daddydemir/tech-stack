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

func DeclareQueue(queueName string) {

	_, err := Channel.QueueDeclare(
		queueName,
		true,
		true,
		true,
		true,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue declare error :", err)
	}
}

func BindQueueToExchange(queueName string) {

	err := Channel.QueueBind(
		queueName,
		"",
		exchange,
		false,
		nil)

	if err != nil {
		log.Println("RabbitMQ queue bind error :", err)
	}
}
