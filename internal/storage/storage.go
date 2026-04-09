package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

type storage struct {
	rdb *redis.Client
	ttl time.Duration
}

func New(ctx context.Context, rdb *redis.Client) (*storage, error) {
	cmd := rdb.Ping(ctx)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", cmd.Err())
	}

	return &storage{
		rdb: rdb,
	}, nil
}

func (s *storage) Put(ctx context.Context, key string, page model.Page) error {
	cmd := s.rdb.Set(ctx, key, page, s.ttl)
	if cmd.Err() != nil {
		return fmt.Errorf("failed to set value %s to the cache: %w", key, cmd.Err())
	}

	return nil
}

func (s *storage) Get(ctx context.Context, key string) (model.Page, error) {
	cmd := s.rdb.Get(ctx, key)
	if cmd.Err() != nil {
		return model.Page{}, fmt.Errorf("failed to get value %s from the cache: %w", key, cmd.Err())
	}

	// TODO: convert value to model.Page
	var _ = cmd
	return model.Page{}, nil
}

func (s *storage) Delete(ctx context.Context, key string) error {
	cmd := s.rdb.Del(ctx, key)
	if cmd.Err() != nil {
		return fmt.Errorf("failed to delete value %s from the cache: %w", key, cmd.Err())
	}

	return nil
}

func (s *storage) Top(ctx context.Context) ([]model.Page, error) {
	// TODO: how to implement it?
	panic("unimplemented")
}
