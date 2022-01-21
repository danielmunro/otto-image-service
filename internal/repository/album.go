package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
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

func (a *AlbumRepository) FindOrCreateProfileAlbumForUser(user *entity.User) *entity.Album {
	album := &entity.Album{}
	a.conn.Where("user_id = ? AND album_type = ?", user.ID, model.PROFILE_PICS).Scan(&album)
	if album.Uuid == nil {
		album = &entity.Album{
			Link:        user.Username,
			AlbumType:   string(model.PROFILE_PICS),
			Name:        user.Username + "'s profile pictures",
			Description: "Profile pictures for " + user.Username,
			User:        user,
			UserID:      user.ID,
			Visibility:  string(model.PUBLIC),
		}
		a.conn.Create(album)
	}
	return album
}
