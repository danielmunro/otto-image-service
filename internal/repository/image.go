package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ImageRepository struct {
	conn *gorm.DB
}

func CreateImageRepository(conn *gorm.DB) *ImageRepository {
	return &ImageRepository{conn}
}

func (i *ImageRepository) Create(image *entity.Image) {
	i.conn.Create(image)
}

func (i *ImageRepository) Update(image *entity.Image) {
	i.conn.Update(image)
}

func (i *ImageRepository) FindByUserAndAlbum(userUuid *uuid.UUID, albumUuid *uuid.UUID) *entity.Image {
	image := &entity.Image{}
	i.conn.Preload("Album").
		Joins("join users on users.id = albums.user_id").
		Where("user_uuid = ? and uuid = ?", userUuid.String(), albumUuid.String()).
		Find(&image)
	return image
}
