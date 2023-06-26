package handlers

import (
	"github.com/dat-guy-defoe/storage/api/handlers/fs"
	"github.com/dat-guy-defoe/storage/api/handlers/minio"
	"github.com/dat-guy-defoe/storage/api/handlers/mongo"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StoreHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	DownloadFile(w http.ResponseWriter, r *http.Request)
	DeleteFile(w http.ResponseWriter, r *http.Request)
}

func BuildHandler() http.Handler {
	router := mux.NewRouter()

	router.Handle("/", &RootHandler{})
	router.Handle("/{storeType}/put", &PutHandler{})
	router.Handle("/{storeType}/get/{fileName}", &GetHandler{})
	router.Handle("/{storeType}/put", &DeleteHandler{})

	return router
}

type RootHandler struct {
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Unrecognized path"))
}

type PutHandler struct {
}

func (h *PutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	switch storeType {
	case "mgo":
		mongo.UploadFile(w, r)
	case "os":
		minio.UploadFile(w, r)
	case "fs":
		fs.UploadFile(w, r)
	}
}

type GetHandler struct {
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	switch storeType {
	case "mgo":
		mongo.DownloadFile(w, r)
	case "os":
		minio.DownloadFile(w, r)
	case "fs":
		fs.DownloadFile(w, r)
	}
}

type DeleteHandler struct {
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	switch storeType {
	case "mgo":
		mongo.DeleteFile(w, r)
	case "os":
		minio.DeleteFile(w, r)
	case "fs":
		fs.DeleteFile(w, r)
	}
}
