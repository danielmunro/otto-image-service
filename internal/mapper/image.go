package mapper

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
)

func GetImageEntityFromNewModel(image *model.NewImage) *entity.Image {
	return &entity.Image{
		Link: image.Link,
	}
}

func GetImageModelFromEntity(image *entity.Image) *model.Image {
	return &model.Image{
		Uuid: image.Uuid.String(),
		Link: image.Link,
		CreatedAt: image.CreatedAt,
	}
}
