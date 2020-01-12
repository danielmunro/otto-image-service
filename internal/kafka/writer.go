package kafka

import (
	"github.com/danielmunro/otto-image-service/internal/constants"
	"github.com/segmentio/kafka-go"
)

func CreateWriter(kafkaHost string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaHost},
		Topic: string(constants.Images),
		Balancer: &kafka.LeastBytes{},
	})
}
