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
	log.Print("find or create album, debug album uuid :: ", albumUuid)
	image := &entity.Image{}
	i.conn.Table("images").
		Joins("join users on users.id = albums.user_id").
		Where("users.uuid = ? and uuid = ?", userUuid.String(), albumUuid.String()).
		Find(&image)
	return image
}
