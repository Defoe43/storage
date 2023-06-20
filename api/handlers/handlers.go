package handlers

import (
	"fmt"
	"github.com/dat-guy-defoe/storage/internal/fs"
	mongodb "github.com/dat-guy-defoe/storage/internal/repo"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var storagePath = os.Getenv("storagePath")

func BuildHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/files/put", uploadFile).Methods("PUT")
	router.HandleFunc("/files/get/{fileName}", downloadFile).Methods("GET")
	router.HandleFunc("/files/delete/{fileName}", deleteFile).Methods("DELETE")

	return router
}

func root(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Unrecognized path"))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		fmt.Printf("uploadFile exec in: %v\n", time.Since(start))
	}(time.Now())

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error retrieving the file")
		log.Println(err)
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)

		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)

		return
	}

	db := mongodb.GetDatabase()

	err = db.PutFile(handler.Filename, fileBytes)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		fmt.Printf("downloadFile exec in: %v\n", time.Since(start))
	}(time.Now())

	vars := mux.Vars(r)
	filename := vars["fileName"]

	db := mongodb.GetDatabase()

	stream, err := db.GetFile(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get the file", http.StatusInternalServerError)

		return
	}
	defer stream.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	_, err = io.Copy(w, stream)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get the file", http.StatusInternalServerError)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		fmt.Printf("deleteFile exec in: %v\n", time.Since(start))
	}(time.Now())

	vars := mux.Vars(r)
	filename := vars["fileName"]

	fileExists := fs.IsFileExist(storagePath + filename)
	if !fileExists {
		log.Printf("File %s does not exist", filename)
		http.Error(w, fmt.Sprintf("File %s does not exist", filename), http.StatusNotFound)

		return
	}

	err := fs.DeleteFile(storagePath + filename)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	} else {
		http.Error(w, fmt.Sprintf("Failed to delete file %s", filename), http.StatusInternalServerError)
	}
}
