package service

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/entity"
	kafka2 "github.com/danielmunro/otto-image-service/internal/kafka"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/google/uuid"
	"log"
	"os"
)

type ImageService struct {
	imageRepository *repository.ImageRepository
	albumRepository *repository.AlbumRepository
	userRepository  *repository.UserRepository
	uploadService   *UploadService
	kafkaWriter     *kafka.Producer
}

func CreateDefaultImageService() *ImageService {
	conn := db.CreateDefaultConnection()
	return CreateImageService(
		repository.CreateImageRepository(conn),
		repository.CreateAlbumRepository(conn),
		repository.CreateUserRepository(conn),
		CreateDefaultUploadService(),
		kafka2.CreateProducer())
}

func CreateImageService(imageRepository *repository.ImageRepository, albumRepository *repository.AlbumRepository, userRepository *repository.UserRepository, uploadService *UploadService, kafkaProducer *kafka.Producer) *ImageService {
	return &ImageService{
		imageRepository,
		albumRepository,
		userRepository,
		uploadService,
		kafkaProducer,
	}
}

func (i *ImageService) CreateNewProfileImage(userUuid uuid.UUID, image *os.File) (*model.Image, error) {
	user, err := i.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		log.Print("no user")
		return nil, err
	}
	album := i.albumRepository.FindOrCreateProfileAlbumForUser(user)
	imageUuid := uuid.New()
	imageEntity := &entity.Image{
		Filename: image.Name(),
		User:     user,
		UserID:   user.ID,
		Album:    album,
		AlbumID:  album.ID,
		Uuid:     &imageUuid,
	}
	s3Key, err := i.uploadService.UploadImage(image)
	if err != nil {
		log.Print("error occurred in image service upload", err)
		return nil, err
	}
	imageEntity.S3Key = s3Key
	i.imageRepository.Create(imageEntity)
	imageModel := mapper.GetImageModelFromEntity(imageEntity)
	data, _ := json.Marshal(imageModel)
	log.Print("publishing image to kafka: ", string(data))
	topic := "images"
	_ = i.kafkaWriter.Produce(
		&kafka.Message{
			Value: data,
			TopicPartition: kafka.TopicPartition{Topic: &topic,
				Partition: kafka.PartitionAny},
		},
		nil)
	image.Close()
	return imageModel, nil
}
