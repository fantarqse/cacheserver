package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"

	"github.com/fantarqse/cacheserver/internal/core/service"
	"github.com/fantarqse/cacheserver/internal/httpserver"
	"github.com/fantarqse/cacheserver/internal/platform/flags"
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

	fs := flags.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", fs.Cache.Address, fs.Cache.Port),
	})
	storage := storage.New(rdb)
	service := service.New(storage)
	server := httpserver.New(service)

	errCh := make(chan error, 1)
	go func() {
		if err := server.Serve(ctx, fs.HTTPServer.Port); err != nil {
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
