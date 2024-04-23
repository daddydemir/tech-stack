package kafka

import (
	"github.com/segmentio/kafka-go"
	"log"
)

func CreateTopic() {

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	configs := []kafka.TopicConfig{
		{
			Topic:             "single-partitioned-topic",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}
	err = conn.CreateTopics(configs...)
	if err != nil {
		log.Println("Error creating: ", err)
	} else {
		log.Println("Successfully created")
	}
}
