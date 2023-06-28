package main

import (
	"github.com/dat-guy-defoe/storage/api/handlers"
	"net/http"
	"os"
)

func main() {
	hs := &http.Server{
		Addr:    ":8080",
		Handler: handlers.BuildHandler(),
	}

	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(-1)
	}
}
