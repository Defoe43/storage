package fs

import (
	"fmt"
	"github.com/dat-guy-defoe/storage/internal/repo/fs"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
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

	fileExists := fs.IsFileExist("./" + handler.Filename)
	if fileExists {
		log.Printf("File %s already exists", handler.Filename)
		http.Error(w, fmt.Sprintf("File %s already exists", handler.Filename), http.StatusBadRequest)

		return
	}

	err = fs.WriteFile("./"+handler.Filename, fileBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["fileName"]

	file, err := fs.GetFile("./" + filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("File %s not found", filename), http.StatusNotFound)

		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)

		return
	}

	fileLength := fmt.Sprintf("%d", fileInfo.Size())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fileLength)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["fileName"]

	fileExists := fs.IsFileExist("./" + filename)
	if !fileExists {
		log.Printf("File %s does not exist", filename)
		http.Error(w, fmt.Sprintf("File %s does not exist", filename), http.StatusNotFound)

		return
	}

	err := fs.DeleteFile(filename)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	} else {
		http.Error(w, fmt.Sprintf("Failed to delete file %s", filename), http.StatusInternalServerError)
	}
}
