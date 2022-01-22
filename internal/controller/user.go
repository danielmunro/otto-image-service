package controller

import (
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/danielmunro/otto-image-service/internal/uuid"
	"net/http"
)

// UploadNewProfilePicV1 - upload a new profile pic
func UploadNewProfilePicV1(w http.ResponseWriter, r *http.Request) {
	userUuid := uuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().
			CreateNewProfileImage(userUuid, tempFile, fileHeader.Filename, fileHeader.Size)
	})
}
