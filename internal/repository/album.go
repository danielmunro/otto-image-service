package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/jinzhu/gorm"
)

type AlbumRepository struct {
	conn *gorm.DB
}

func CreateAlbumRepository(conn *gorm.DB) *AlbumRepository {
	return &AlbumRepository{conn}
}

func (a *AlbumRepository) Create(album *entity.Album) {
	a.conn.Create(album)
}
