package controller

import (
	"github.com/danielmunro/otto-image-service/internal/auth/model"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

// UploadNewProfileImageV1 - upload a new profile pic
func UploadNewProfileImageV1(w http.ResponseWriter, r *http.Request) {
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().
			CreateNewProfileImage(uuid.MustParse(session.User.Uuid), tempFile, fileHeader.Filename, fileHeader.Size)
	})
}
