package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"

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
			- address, password, etc.
			- max count of items
			- size of one item in bytes
			- size of all items in bytes
			- TTL
		- logger
	*/

	port := flag.Int("port", 8080, "a port of an http server")
	redisAddr := flag.String("redis-addr", "locahlhost", "an address of redis")
	redisPort := flag.Int("redis-port", 6379, "an address of redis")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", *redisAddr, *redisPort),
	})
	storage := storage.New(rdb)
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
