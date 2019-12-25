package mapper

import (
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/google/uuid"
)

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	return &entity.User{
		Uuid:       &userUuid,
		Username:   user.Username,
	}
}
