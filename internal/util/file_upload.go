package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

func StoreUploadedFile(r *http.Request) (*os.File, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("/tmp", "upload-*")
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
