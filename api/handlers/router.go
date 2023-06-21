package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func BuildHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/files/put", uploadFile).Methods("PUT")
	router.HandleFunc("/os/put", minioUploadFile).Methods("PUT")
	router.HandleFunc("/files/get/{fileName}", downloadFile).Methods("GET")
	router.HandleFunc("/os/get/{fileName}", minioDownloadFile).Methods("GET")
	router.HandleFunc("/files/delete/{fileName}", deleteFile).Methods("DELETE")

	return router
}
