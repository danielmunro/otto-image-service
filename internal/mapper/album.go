package mapper

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
)

func GetAlbumEntityFromNewModel(album *model.NewAlbum) *entity.Album {
	return &entity.Album{
		Link:        album.Link,
		AlbumType:   string(model.UserCreated),
		Name:        album.Name,
		Description: album.Description,
		Visibility:  string(album.Visibility),
	}
}

func GetAlbumModelFromEntity(album *entity.Album) *model.Album {
	return &model.Album{
		Uuid:        album.Uuid.String(),
		Link:        album.Link,
		Name:        album.Name,
		Description: album.Description,
		Visibility:  model.Visibility(album.Visibility),
		AlbumType:   model.AlbumType(album.AlbumType),
		CreatedAt:   album.CreatedAt,
		UpdatedAt:   album.UpdatedAt,
	}
}
