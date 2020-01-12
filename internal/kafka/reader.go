package kafka

import (
	"github.com/danielmunro/otto-image-service/internal/constants"
	"github.com/segmentio/kafka-go"
)

func GetReader(broker string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     string(constants.Users),
		GroupID: "image_service",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return r
}
