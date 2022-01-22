package repository

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/jinzhu/gorm"
	"log"
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
	log.Print("find or create profile album, user :: ", user.ID)
	album := &entity.Album{}
	albumType := string(model.PROFILE_PICS)
	a.conn.Where("user_id = ? AND album_type = ?", user.ID, albumType).Scan(&album)
	if album.Uuid == nil {
		album = &entity.Album{
			Link:        user.Username,
			AlbumType:   albumType,
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
