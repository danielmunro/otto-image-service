package service

import (
	"github.com/danielmunro/otto-image-service/internal/db"
	"github.com/danielmunro/otto-image-service/internal/mapper"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/repository"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/danielmunro/otto-image-service/internal/test"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"testing"
)

func createTestUser() *model.User {
	userEntity := test.CreateTestUser()
	repo := repository.CreateUserRepository(db.CreateDefaultConnection())
	repo.Create(userEntity)
	userModel := mapper.GetUserModelFromEntity(userEntity)
	return &userModel
}

func createFile() (multipart.File, error) {
	filePath := "sample.jpeg"
	return os.Open(filePath)
}

func Test_Can_UploadImage(t *testing.T) {
	// setup
	user := createTestUser()

	// given
	filename := "sample.jpeg"
	file, _ := createFile()
	fstat, _ := os.Stat(filename)

	// when
	imageModel, err := service.CreateDefaultImageService().
		CreateNewProfileImage(uuid.MustParse(user.Uuid), file, filename, fstat.Size())

	// then
	test.Assert(t, imageModel != nil)
	test.Assert(t, err == nil)
}
