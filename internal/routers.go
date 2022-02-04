/*
 * Otto image service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package internal

import (
	"fmt"
	"github.com/danielmunro/otto-image-service/internal/controller"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreateNewAlbumV1",
		strings.ToUpper("Post"),
		"/album",
		controller.CreateNewAlbumV1,
	},

	{
		"CreateNewImageV1",
		strings.ToUpper("Post"),
		"/album/{link}/image",
		controller.CreateNewImageV1,
	},

	{
		"GetAlbumV1",
		strings.ToUpper("Get"),
		"/album/{uuid}",
		controller.GetAlbumV1,
	},

	{
		"GetAlbumsForUserV1",
		strings.ToUpper("Get"),
		"/user/{username}/album",
		controller.GetAlbumsForUserV1,
	},

	{
		"GetImageV1",
		strings.ToUpper("Get"),
		"/image/{uuid}",
		controller.GetImageV1,
	},

	{
		"GetImagesForAlbumV1",
		strings.ToUpper("Get"),
		"/album/{uuid}/image",
		controller.GetImagesForAlbumV1,
	},

	{
		"UploadNewLivestreamImageV1",
		strings.ToUpper("Post"),
		"/album/livestream",
		controller.UploadNewLivestreamImageV1,
	},

	{
		"UploadNewProfileImageV1",
		strings.ToUpper("Post"),
		"/album/profile",
		controller.UploadNewProfileImageV1,
	},
}
