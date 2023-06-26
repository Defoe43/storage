package handlers

import (
	"fmt"
	"github.com/dat-guy-defoe/storage/api/handlers/fs"
	"github.com/dat-guy-defoe/storage/api/handlers/minio"
	"github.com/dat-guy-defoe/storage/api/handlers/mongo"
	"github.com/gorilla/mux"
	"net/http"
)

type PutHandler struct {
	factory StorageHandlerFactory
}

func (h *PutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	handler, err := h.factory.Create(storeType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler.UploadFile(w, r)
}

type GetHandler struct {
	factory StorageHandlerFactory
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	handler, err := h.factory.Create(storeType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler.DownloadFile(w, r)
}

type DeleteHandler struct {
	factory StorageHandlerFactory
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storeType := mux.Vars(r)["storeType"]

	handler, err := h.factory.Create(storeType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler.DeleteFile(w, r)
}

type StoreHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	DownloadFile(w http.ResponseWriter, r *http.Request)
	DeleteFile(w http.ResponseWriter, r *http.Request)
}

type StorageHandlerFactory struct {
}

func (f *StorageHandlerFactory) Create(storeType string) (StoreHandler, error) {
	switch storeType {
	case "mgo":
		return &mongo.Handler{}, nil
	case "os":
		return &minio.Handler{}, nil
	case "fs":
		return &fs.Handler{}, nil
	default:
		return nil, fmt.Errorf("invalid store type: %s", storeType)
	}
}

type RootHandler struct {
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Unrecognized path"))
}
