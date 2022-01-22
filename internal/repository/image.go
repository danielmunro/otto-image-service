package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"log"
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

func (i *ImageRepository) FindByUserAndAlbum(userUuid *uuid.UUID, albumUuid *uuid.UUID) *entity.Image {
	log.Print("find or create album, debug user uuid :: ", userUuid)
	log.Print("album uuid :: ", albumUuid)
	image := &entity.Image{}
	i.conn.Preload("User").
		Table("images").
		Joins("join albums on albums.id = images.album_id").
		Joins("join users on users.id = images.user_id").
		Where("albums.uuid = ? and users.uuid = ?", albumUuid, userUuid).
		First(&image)
	return image
}
