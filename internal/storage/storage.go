package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

type storage struct {
	rdb *redis.Client
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

func (s *storage) Put(ctx context.Context, key string, page model.Page, ttl time.Duration) error {
	if _, err := s.rdb.Set(ctx, key, page, ttl).Result(); err != nil {
		return fmt.Errorf("failed to set value %s to the cache: %w", key, err)
	}

	return nil
}

func (s *storage) Get(ctx context.Context, key string) (model.Page, error) {
	var page model.Page

	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return model.Page{}, fmt.Errorf("failed to get value %s from the cache: %w", key, err)
	}

	if err := json.Unmarshal([]byte(val), &page); err != nil {
		return model.Page{}, fmt.Errorf("failed to unmarshal value %s: %w", key, err)
	}

	return page, nil
}

func (s *storage) Delete(ctx context.Context, key string) error {
	if _, err := s.rdb.Del(ctx, key).Result(); err != nil {
		return fmt.Errorf("failed to delete value %s from the cache: %w", key, err)
	}

	return nil
}

func (s *storage) Top(ctx context.Context) ([]model.Page, error) {
	var pages []model.Page
	if err := s.rdb.HGetAll(ctx, "").Scan(&pages); err != nil {
		return nil, fmt.Errorf("failed to get all records: %w", err)
	}

	return pages, nil
}
