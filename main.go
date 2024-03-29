/*
 * Otto image service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/danielmunro/otto-image-service/internal"
	"github.com/danielmunro/otto-image-service/internal/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	serveHttp()
}
func serveHttp() {
	router := internal.NewRouter()
	handler := cors.AllowAll().Handler(router)
	log.Print("listening on 8082")
	log.Fatal(http.ListenAndServe(":8082",
		middleware.FileSizeLimit(middleware.ContentTypeMiddleware(handler))))
}
