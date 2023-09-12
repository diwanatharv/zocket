package dataaccess

import (
	"github.com/IBM/sarama"
	"log"
)

func ProduceToKafka(producer sarama.SyncProducer, topic string, key string, value string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	_, _, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}

	log.Printf("Message sent to topic %s", topic)
	return nil
}
