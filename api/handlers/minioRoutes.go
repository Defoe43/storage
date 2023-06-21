package handlers

import (
	"bytes"
	"fmt"
	"github.com/dat-guy-defoe/storage/internal/repo/minio"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

func minioUploadFile(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		fmt.Printf("minioUploadFile exec in: %v\n", time.Since(start))
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

	objSt, err := minio.GetClient()
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)

		return
	}
	bytesReader := bytes.NewReader(fileBytes)
	reader := io.Reader(bytesReader)
	err = objSt.PutObject(handler.Filename, &reader)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func minioDownloadFile(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		fmt.Printf("downloadFile exec in: %v\n", time.Since(start))
	}(time.Now())

	vars := mux.Vars(r)
	filename := vars["fileName"]

	objSt, err := minio.GetClient()
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusInternalServerError)

		return
	}

	data, err := objSt.GetObject(filename)
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusInternalServerError)
	}
}
