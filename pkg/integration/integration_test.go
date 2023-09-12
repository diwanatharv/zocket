package integration

import (
	"awesomeProject6/pkg/config/kafka"
	"awesomeProject6/pkg/dataaccess"
	"strconv"
	"testing"
)

func TestKafkaIntegration(t *testing.T) {
	// Initialize Kafka producer and consumer
	producer, err := kafka.CreateKafkaproducer()
	if err != nil {
		t.Fatalf("Failed to create Kafka producer: %v", err)
	}
	consumer, err := kafka.CreateKafkaConsumer()
	if err != nil {
		t.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Assuming you have a valid product ID for testing
	productID := 1

	// Produce a Kafka message
	topic := "product-topic"
	key := "product-created"
	value := strconv.Itoa(productID)

	err = dataaccess.ProduceToKafka(producer, topic, key, value)
	if err != nil {
		t.Fatalf("Failed to produce message to Kafka: %v", err)
	}

	// Consume the Kafka message and test further interactions
	err = dataaccess.Consumemessage(topic, consumer)
	if err != nil {
		t.Fatalf("Failed to consume message from Kafka: %v", err)
	}
}
