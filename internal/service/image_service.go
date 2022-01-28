package service

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/entity"
	kafka2 "github.com/danielmunro/otto-image-service/internal/kafka"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
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

func (i *ImageService) CreateNewLivestreamImage(userUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	user, err := i.userRepository.FindOneByUuid(userUuid.String())
	if user.Uuid == nil || err != nil {
		log.Print("error finding user :: ", err)
		return
	}
	album := i.albumRepository.FindOrCreateLivestreamAlbumForUser(user)
	s3Key, err := i.uploadService.UploadImage(file, filename, filesize)
	if err != nil {
		log.Print("error occurred in image service upload", err)
		return
	}
	imageEntity := i.createNewLivestreamImageEntity(user, album)
	imageEntity.S3Key = s3Key
	i.imageRepository.Create(imageEntity)
	imageModel = mapper.GetImageModelFromEntity(imageEntity)
	return
}

func (i *ImageService) CreateNewProfileImage(userUuid uuid.UUID, file multipart.File, filename string, filesize int64) (imageModel *model.Image, err error) {
	user, err := i.userRepository.FindOneByUuid(userUuid.String())
	if user.Uuid == nil || err != nil {
		log.Print("error finding user :: ", err)
		return
	}
	album := i.albumRepository.FindOrCreateLivestreamAlbumForUser(user)
	s3Key, err := i.uploadService.UploadImage(file, filename, filesize)
	if err != nil {
		log.Print("error occurred in image service upload", err)
		return
	}
	imageEntity := i.createNewLivestreamImageEntity(user, album)
	imageEntity.S3Key = s3Key
	i.imageRepository.Save(imageEntity)
	imageModel = mapper.GetImageModelFromEntity(imageEntity)
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
	return
}

func (i *ImageService) GetImage(imageUuid uuid.UUID) (*model.Image, error) {
	imageEntity := i.imageRepository.FindByUuid(&imageUuid)
	if imageEntity.ID == 0 {
		return nil, errors.New("image not found")
	}
	return mapper.GetImageModelFromEntity(imageEntity), nil
}

func (i *ImageService) findOrCreateProfileImage(user *entity.User, album *entity.Album) (imageEntity *entity.Image) {
	imageEntity = i.imageRepository.FindByUserAndAlbum(user.Uuid, album.Uuid)
	if imageEntity.Uuid == nil {
		log.Print("profile pic not found, creating new one, user :: ", user.Uuid)
		imageEntity = &entity.Image{
			Filename: "",
			User:     user,
			UserID:   user.ID,
			Album:    album,
			AlbumID:  album.ID,
		}
		i.imageRepository.Create(imageEntity)
	}
	return
}

func (i *ImageService) createNewLivestreamImageEntity(user *entity.User, album *entity.Album) *entity.Image {
	return &entity.Image{
		Filename: "",
		User:     user,
		UserID:   user.ID,
		Album:    album,
		AlbumID:  album.ID,
	}
}
