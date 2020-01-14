package service

import (
	"context"
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/entity"
	kafka2 "github.com/danielmunro/otto-image-service/internal/kafka"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

type ImageService struct {
	imageRepository *repository.ImageRepository
	albumRepository *repository.AlbumRepository
	userRepository *repository.UserRepository
	uploadService *UploadService
	kafkaWriter *kafka.Writer
}

func CreateDefaultImageService() *ImageService {
	conn := db.CreateDefaultConnection()
	return CreateImageService(
		repository.CreateImageRepository(conn),
		repository.CreateAlbumRepository(conn),
		repository.CreateUserRepository(conn),
		CreateDefaultUploadService(),
		kafka2.CreateWriter(os.Getenv("KAFKA_HOST")))
}

func CreateImageService(imageRepository *repository.ImageRepository, albumRepository *repository.AlbumRepository, userRepository *repository.UserRepository, uploadService *UploadService, kafkaWriter *kafka.Writer) *ImageService {
	return &ImageService{
		imageRepository,
		albumRepository,
		userRepository,
		uploadService,
		kafkaWriter,
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
	imageEntity.S3Key = s3Key.String()
	i.imageRepository.Create(imageEntity)
	imageModel := mapper.GetImageModelFromEntity(imageEntity)
	data, _ := json.Marshal(imageModel)
	log.Print("publishing image to kafka: ", string(data))
	_ = i.kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: data})
	image.Close()
	return imageModel, nil
}