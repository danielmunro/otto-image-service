package controller

import "net/http"

// CreateNewImageV1 - create a new image
func CreateNewImageV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
