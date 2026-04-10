package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/fantarqse/cacheserver/internal/core/model"
	"github.com/fantarqse/cacheserver/internal/core/service"
)

// It must be configurable value, not a constant.
const basicKey string = "cacheserver"

type storage struct {
	rdb *redis.Client
}

func New(ctx context.Context, rdb *redis.Client) (*storage, error) {
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &storage{
		rdb: rdb,
	}, nil
}

func (s *storage) Put(ctx context.Context, key string, page model.Page, ttl time.Duration) error {
	key = fmt.Sprintf("%s:%s", basicKey, key)

	if _, err := s.rdb.Set(ctx, key, page, ttl).Result(); err != nil {
		return fmt.Errorf("failed to set value %s to the cache: %w", key, err)
	}

	return nil
}

func (s *storage) Get(ctx context.Context, key string) (model.Page, error) {
	var (
		page    model.Page
		fullKey = fmt.Sprintf("%s:%s", basicKey, key)
	)

	val, err := s.rdb.Get(ctx, fullKey).Result()
	switch err {
	case nil:
		// Do nothing.
	case redis.Nil:
		return model.Page{}, service.ErrNotFound
	default:
		return model.Page{}, fmt.Errorf("failed to get value %s from the cache: %w", fullKey, err)
	}

	if err := json.Unmarshal([]byte(val), &page); err != nil {
		return model.Page{}, fmt.Errorf("failed to unmarshal value %s: %w", fullKey, err)
	}

	return page, nil
}

func (s *storage) Delete(ctx context.Context, key string) error {
	key = fmt.Sprintf("%s:%s", basicKey, key)

	if _, err := s.rdb.Del(ctx, key).Result(); err != nil {
		return fmt.Errorf("failed to delete value %s from the cache: %w", key, err)
	}

	return nil
}

func (s *storage) Top(ctx context.Context) ([]model.Page, error) {
	var pages []model.Page

	if err := s.rdb.HGetAll(ctx, basicKey).Scan(&pages); err != nil {
		return nil, fmt.Errorf("failed to get all records: %w", err)
	}

	return pages, nil
}
