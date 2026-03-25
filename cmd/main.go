package main

import (
	"context"
	"log"
	"os"

	"github.com/fantarqse/cacheserver/internal/core/service"
	"github.com/fantarqse/cacheserver/internal/httpserver"
	"github.com/fantarqse/cacheserver/internal/storage"
)

func main() {
	// TODO: logger?
	// Do I want to use a standard slog or zap?

	// TODO: config?
	// I am considering using flags to configure the entire app.

	storage := storage.New()
	service := service.New(storage)
	server := httpserver.New(service)

	if err := server.Serve(context.Background()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
