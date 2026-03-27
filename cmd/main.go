package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/fantarqse/cacheserver/internal/core/service"
	"github.com/fantarqse/cacheserver/internal/httpserver"
	"github.com/fantarqse/cacheserver/internal/storage"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	// TODO: logger?
	// Do I want to use a standard slog or zap?

	// TODO: config?
	// I am considering using flags to configure the entire app.
	/*
		- http server
			- port
		- storage
			- max count of items
			- size of one item in bytes
			- size of all items in bytes
			- TTL
		- logger
	*/

	storage := storage.New()
	service := service.New(storage)
	server := httpserver.New(service)

	go func() {
		if err := server.Serve(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	return nil
}
