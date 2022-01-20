package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func StoreUploadedFile(r *http.Request) (*os.File, error) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Print("error accessing form image :: ", err)
		return nil, err
	}
	defer file.Close()
	ext := filepath.Ext(handler.Filename)
	tempFile, err := ioutil.TempFile("/tmp", "upload-*."+ext)
	if err != nil {
		return nil, err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	_, _ = tempFile.Write(fileBytes)
	return tempFile, nil
}
