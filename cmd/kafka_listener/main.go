package main

import (
	"context"
	"github.com/danielmunro/otto-image-service/internal/constants"
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load()
	log.Print("connecting to kafka: ", os.Getenv("KAFKA_HOST"))
	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_HOST"), string(constants.Users), 0)
	if err != nil {
		log.Print("fail")
		log.Fatal(err)
	}
	_ = conn.SetReadDeadline(time.Now().Add(10*time.Second))
	for {
		batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
		userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
		err := ParseBatch(userRepository, batch)
		if err != nil {
			break
		}
		_ = batch.Close()
	}
	_ = conn.Close()
}

func ParseBatch(userRepository *repository.UserRepository, batch *kafka.Batch) error {
	b := make([]byte, 10e3)            // 10KB max per message
	for {
		readLen, err := batch.Read(b)
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Print("error received", err)
			return err
		}
		data := b[:readLen]
		userModel, err := model.DecodeMessageToUser(data)
		if err != nil {
			log.Print("error decoding message to user, skipping")
			continue
		}
		userEntity, err := userRepository.FindOneByUuid(userModel.Uuid)
		if err == nil {
			log.Print("skip user add")
		} else {
			log.Print("create")
			userEntity = mapper.GetUserEntityFromModel(userModel)
			userRepository.Create(userEntity)
		}
	}
	return nil
}
