package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

// CreateNewAlbumV1 - create a new album
func CreateNewAlbumV1(w http.ResponseWriter, r *http.Request) {
	newAlbum := model.DecodeRequestToNewAlbum(r)
	service.CreateDefaultAuthService().DoWithValidSessionAndUser(w, r, uuid.MustParse(newAlbum.User.Uuid), func () (interface{}, error) {
		album := service.CreateDefaultAlbumService().CreateAlbum(newAlbum)
		data, err := json.Marshal(album)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}
		return data, err
	})
}
