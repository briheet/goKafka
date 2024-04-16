package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func newKafkaProducer() (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func main() {
	producer, err := newKafkaProducer()
	if err != nil {
		fmt.Printf("Failed to create Kafka producer: %s\n", err)
		return
	}
	defer producer.Close()

	topic := "your-topic"
	message := "Hello, Kafka!"

	err = produceMessage(producer, topic, message)
	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return
	}
}

func produceMessage(producer *kafka.Producer, topic string, message string) error {
	deliveryChan := make(chan kafka.Event)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	fmt.Printf("Message '%s' delivered to topic %s [%d] at offset %v\n",
		message, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	close(deliveryChan)

	return nil
}
