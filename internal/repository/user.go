package repository

import (
	"errors"
	"github.com/danielmunro/otto-image-service/internal/constants"
	"github.com/danielmunro/otto-image-service/internal/entity"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func CreateUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (u *UserRepository) FindOneByUuid(uuid string) (*entity.User, error) {
	user := &entity.User{}
	u.conn.Where("uuid = ?", uuid).Find(user)
	if user.ID == 0 {
		return nil, errors.New(constants.ErrorMessageUserNotFound)
	}
	return user, nil
}

func (u *UserRepository) Create(user *entity.User) {
	u.conn.Create(user)
}

func (u *UserRepository) Update(user *entity.User) {
	u.conn.Model(&user).Updates(user)
}

func (u *UserRepository) Delete(user *entity.User) {
	u.conn.Delete(user)
}
