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

func (i *ImageRepository) Save(image *entity.Image) {
	i.conn.Save(image)
}

func (i *ImageRepository) FindByUuid(imageUuid *uuid.UUID) *entity.Image {
	image := &entity.Image{}
	i.conn.Preload("User").
		Table("images").
		Where("uuid = ?", imageUuid).
		First(&image)
	return image
}

func (i *ImageRepository) FindByUserAndAlbum(userUuid *uuid.UUID, albumUuid *uuid.UUID) *entity.Image {
	image := &entity.Image{}
	i.conn.Preload("User").
		Table("images").
		Joins("join albums on albums.id = images.album_id").
		Joins("join users on users.id = images.user_id").
		Where("albums.uuid = ? and users.uuid = ?", albumUuid, userUuid).
		First(&image)
	return image
}
