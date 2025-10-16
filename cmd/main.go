package main

import (
	"log"

	"github.com/fantarqse/cacheserver/internal/core/service"
	"github.com/fantarqse/cacheserver/internal/httpserver"
	"github.com/fantarqse/cacheserver/internal/storage"
)

func main() {
	// TODO: logger?
	// TODO: config

	storage := storage.New()
	service := service.New(storage)
	server := httpserver.New(service)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
