package main

import (
	"context"
	"flag"
	"fmt"
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

	port := flag.Int("port", 8080, "a port of an http server")
	flag.Parse()

	storage := storage.New()
	service := service.New(storage)
	server := httpserver.New(service)

	errCh := make(chan error, 1)
	go func() {
		if err := server.Serve(ctx, *port); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		// process further
		log.Println("shutting down")
	case err := <-errCh:
		return fmt.Errorf("failed to start the server: %w", err)
	}

	return nil
}
