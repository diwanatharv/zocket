package kafka

import (
	"github.com/IBM/sarama"
	"github.com/labstack/gommon/log"
)

type Config struct {
	KafkaBootstrapServers string
}

// LoadConfig loads the application configuration from environment variables or a configuration file.
func LoadConfig() *Config {
	return &Config{
		KafkaBootstrapServers: "localhost:9092",
	}
}

func CreateKafkaproducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // This line was corrected
	manager := LoadConfig()
	producer, err := sarama.NewSyncProducer([]string{manager.KafkaBootstrapServers}, config)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return producer, nil
}

func CreateKafkaConsumer() (consumer sarama.Consumer, err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true // This line was corrected
	manager := LoadConfig()
	consumer, err = sarama.NewConsumer([]string{manager.KafkaBootstrapServers}, config)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return consumer, nil
}
