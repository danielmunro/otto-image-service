package service

import (
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/google/uuid"
	"log"
	"os"
)

type ImageService struct {
	imageRepository *repository.ImageRepository
	userRepository *repository.UserRepository
	uploadService *UploadService
}

func CreateDefaultImageService() *ImageService {
	conn := db.CreateDefaultConnection()
	return CreateImageService(
		repository.CreateImageRepository(conn),
		repository.CreateUserRepository(conn),
		CreateDefaultUploadService())
}

func CreateImageService(imageRepository *repository.ImageRepository, userRepository *repository.UserRepository, uploadService *UploadService) *ImageService {
	return &ImageService{
		imageRepository,
		userRepository,
		uploadService,
	}
}

func (i *ImageService) CreateImage(userUuid uuid.UUID, image *os.File) *model.Image {
	// imageEntity := mapper.GetImageEntityFromNewModel(image)
	// i.imageRepository.Create(imageEntity)
	// return mapper.GetImageModelFromEntity(imageEntity)
	user, err := i.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		log.Print("no user")
		return nil
	}
	imageUuid := uuid.New()
	imageEntity := &entity.Image{
		Filename: image.Name(),
		User:     user,
		UserID:   user.ID,
		Album:    nil,
		AlbumID:  0,
		Uuid:     &imageUuid,
	}
	s3Key := i.uploadService.UploadImage(image)
	imageEntity.S3Key = s3Key.String()
	return mapper.GetImageModelFromEntity(imageEntity)
}