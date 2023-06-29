package main

import (
	"github.com/dat-guy-defoe/storage/api/handlers"
	"github.com/dat-guy-defoe/storage/internal/repo/mongo"
	"log"
	"net/http"
	"os"
)

func main() {
	mongoRepo := mongo.Repository{}
	err := mongoRepo.Connect()
	if err != nil {
		log.Fatalf("Cannot connect to the MongoDB: %s", err)
	}

	hs := &http.Server{
		Addr:    ":8080",
		Handler: handlers.BuildHandler(),
	}

	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(-1)
	}
}
