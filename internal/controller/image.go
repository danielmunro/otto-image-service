package controller

import (
	"github.com/danielmunro/otto-image-service/internal/auth/model"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

// CreateNewImageV1 - create a new image
func CreateNewImageV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// UploadNewLivestreamImageV1 - upload a new image
func UploadNewLivestreamImageV1(w http.ResponseWriter, r *http.Request) {
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		tempFile, fileHeader, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		return service.CreateDefaultImageService().CreateNewLivestreamImage(
			uuid.MustParse(session.User.Uuid),
			tempFile,
			fileHeader.Filename,
			fileHeader.Size,
		)
	})
}
