package mapper

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
)

func GetImageModelFromEntity(image *entity.Image) *model.Image {
	return &model.Image{
		Uuid:      image.Uuid.String(),
		Link:      image.Link,
		S3Key:     image.S3Key,
		CreatedAt: image.CreatedAt,
		User:      GetUserModelFromEntity(image.User),
	}
}
