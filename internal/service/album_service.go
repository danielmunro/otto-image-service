package service

import (
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
)

type AlbumService struct {
	userRepository *repository.UserRepository
	albumRepository *repository.AlbumRepository
}

func CreateDefaultAlbumService() *AlbumService {
	conn := db.CreateDefaultConnection()
	return CreateAlbumService(
		repository.CreateAlbumRepository(conn),
		repository.CreateUserRepository(conn))
}

func CreateAlbumService(albumRepository *repository.AlbumRepository, userRepository *repository.UserRepository) *AlbumService {
	return &AlbumService{
		userRepository,
		albumRepository,
	}
}

func (a *AlbumService) CreateAlbum(album *model.NewAlbum) *model.Album {
	albumEntity := mapper.GetAlbumEntityFromNewModel(album)
	a.albumRepository.Create(albumEntity)
	return mapper.GetAlbumModelFromEntity(albumEntity)
}