package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/auth/model"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

// GetImageV1 - get an image
func GetImageV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := uuid.MustParse(params["uuid"])
	imageModel, err := service.CreateDefaultImageService().GetImage(uuidParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(imageModel)
	_, _ = w.Write(data)
}
