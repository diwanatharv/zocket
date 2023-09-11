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
	configure := sarama.NewConfig()
	configure.Producer.Return.Successes = true
	manager := LoadConfig()
	producer, err := sarama.NewSyncProducer([]string{manager.KafkaBootstrapServers}, configure)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return producer, nil
}
func CreateKafkaConsumer() (consumer sarama.Consumer, err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	manager := LoadConfig()
	consumers, err := sarama.NewConsumer([]string{manager.KafkaBootstrapServers}, config)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return consumers, nil
}
