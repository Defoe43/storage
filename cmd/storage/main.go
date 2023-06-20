package main

import (
	"fmt"
	"github.com/dat-guy-defoe/storage/api/handlers"
	mongodb "github.com/dat-guy-defoe/storage/internal/repo/mongo"
	"log"
	"net/http"
	"os"
	"time"
)

var serverPort = os.Getenv("serverPort")

func main() {
	_, err := mongodb.GetMongoConnection("mongodb://localhost:27017", "local", 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to connect to the MongoDB: %v", err)
	}

	hs := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: handlers.BuildHandler(),
	}

	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(-1)
	}
}
