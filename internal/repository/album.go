package repository

import (
	"errors"
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/google/uuid"
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
	return a.FindOrCreateAlbumByType(user, string(model.ProfilePics))
}

func (a *AlbumRepository) FindOrCreateLivestreamAlbumForUser(user *entity.User) *entity.Album {
	return a.FindOrCreateAlbumByType(user, string(model.Livestream))
}

func (a *AlbumRepository) FindOrCreateAlbumByType(user *entity.User, albumType string) *entity.Album {
	log.Print("find or create profile album, user :: ", user.ID)
	album := &entity.Album{}
	a.conn.
		Table("albums").
		Where("user_id = ? AND album_type = ?", user.ID, albumType).
		Scan(&album)
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

func (a *AlbumRepository) FindAllByUser(userEntity *entity.User) []*entity.Album {
	var albumEntities []*entity.Album
	a.conn.Preload("User").
		Table("albums").
		Where("user_id = ?", userEntity.ID).
		Find(&albumEntities)
	return albumEntities
}

func (a *AlbumRepository) FindOne(albumUuid uuid.UUID) (*entity.Album, error) {
	albumEntity := &entity.Album{}
	a.conn.Preload("User").
		Preload("Images").
		Table("albums").
		Where("uuid = ?", albumUuid).
		Order("images.id").
		Find(albumEntity)
	if albumEntity.ID == 0 {
		return nil, errors.New("album not found")
	}
	return albumEntity, nil
}
