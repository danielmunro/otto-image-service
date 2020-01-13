package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/danielmunro/otto-image-service/internal/util"
	"github.com/danielmunro/otto-image-service/internal/uuid"
	"net/http"
)

// UploadNewProfilePicV1 - upload a new profile pic
func UploadNewProfilePicV1(w http.ResponseWriter, r *http.Request) {
	userUuid := uuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().DoWithValidSessionAndUser(w, r, userUuid, func () (interface{}, error) {
		tempFile, err := util.StoreUploadedFile(r)
		if err != nil {
			return nil, err
		}
		image, err := service.CreateDefaultImageService().CreateNewProfileImage(userUuid, tempFile)
		if err != nil {
			return nil, err
		}
		data, err := json.Marshal(image)
		if err == nil {
			_, _ = w.Write(data)
		}
		return data, err
	})
}
