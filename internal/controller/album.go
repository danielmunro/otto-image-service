package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/model"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateNewAlbumV1 - create a new album
func CreateNewAlbumV1(w http.ResponseWriter, r *http.Request) {
	newAlbum := model.DecodeRequestToNewAlbum(r)
	service.CreateDefaultAuthService().DoWithValidSessionAndUser(w, r, uuid.MustParse(newAlbum.User.Uuid), func() (interface{}, error) {
		album := service.CreateDefaultAlbumService().CreateAlbum(newAlbum)
		data, err := json.Marshal(album)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}
		return data, err
	})
}

// GetAlbumV1 - create a new album
func GetAlbumV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := params["uuid"]
	album, err := service.CreateDefaultAlbumService().GetAlbum(uuid.MustParse(uuidParam))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(album)
	_, _ = w.Write(data)
}

// GetAlbumsForUserV1 - create a new album
func GetAlbumsForUserV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := params["uuid"]
	albums, err := service.CreateDefaultAlbumService().GetAlbumsForUser(uuid.MustParse(uuidParam))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(albums)
	_, _ = w.Write(data)
}
