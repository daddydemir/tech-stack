package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Channel *amqp.Channel

func ConnectToRabbitMQ() {
	connectionString := "amqp://rabbit-user:password@localhost:5672"
	dial, err := amqp.Dial(connectionString)
	if err != nil {
		log.Println("RabbitMQ connection error :", err)
	} else {
		log.Println("Connecting to RabbitMQ")
	}
	Channel, err = dial.Channel()
	if err != nil {
		log.Println("RabbitMQ channel error :", err)
	}
}
