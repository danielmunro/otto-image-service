package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/jinzhu/gorm"
)

type ImageRepository struct {
	conn *gorm.DB
}

func CreateImageRepository(conn *gorm.DB) *ImageRepository {
	return &ImageRepository{conn}
}

func (a *ImageRepository) Create(image *entity.Image) {
	a.conn.Create(image)
}
