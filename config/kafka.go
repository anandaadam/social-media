package config

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func KafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
}
