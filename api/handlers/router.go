package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func BuildHandler() http.Handler {
	router := mux.NewRouter()

	router.Handle("/", &RootHandler{})
	router.Handle("/{storeType}/put", &PutHandler{})
	router.Handle("/{storeType}/get/{fileName}", &GetHandler{})
	router.Handle("/{storeType}/delete/{fileName}", &DeleteHandler{})

	return router
}
