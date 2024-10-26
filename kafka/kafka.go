// kafka/kafka.go
package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer

func InitProducer() {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatal("Failed to create Kafka producer:", err)
	}

	log.Println("Kafka producer created")
}

func Publish(topic string, message string) error {
	return Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
}
