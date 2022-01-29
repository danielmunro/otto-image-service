package service

import (
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/google/uuid"
	"log"
)

type AlbumService struct {
	userRepository  *repository.UserRepository
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

func (a *AlbumService) GetAlbum(albumUuid uuid.UUID) (*model.Album, error) {
	log.Print("GetAlbum with uuid :: ", albumUuid)
	albumEntity, err := a.albumRepository.FindOne(albumUuid)
	if err != nil {
		return nil, err
	}
	return mapper.GetAlbumModelFromEntity(albumEntity), nil
}

func (a *AlbumService) GetAlbumsForUser(username string) ([]*model.Album, error) {
	userEntity, err := a.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	albumEntities := a.albumRepository.FindAllByUser(userEntity)
	return mapper.GetAlbumModelsFromEntities(albumEntities), nil
}
