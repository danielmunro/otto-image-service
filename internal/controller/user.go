package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-image-service/internal/service"
	"github.com/danielmunro/otto-image-service/internal/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

// UploadNewProfilePicV1 - upload a new profile pic
func UploadNewProfilePicV1(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("image")
	if err != nil {
		log.Print("error getting file from post", err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("/tmp", "upload-*")
	if err != nil {
		log.Print("error writing temp file", err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("error reading file", err)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// newImage := model.DecodeRequestToNewImage(r)
	userUuid := uuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().DoWithValidSessionAndUser(w, r, userUuid, func () (interface{}, error) {
		image := service.CreateDefaultImageService().CreateImage(userUuid, tempFile)
		data, err := json.Marshal(image)
		if err == nil {
			_, _ = w.Write(data)
		}
		return data, err
	})
}
