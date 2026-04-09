package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

type storage struct {
	rdb *redis.Client
}

func New(ctx context.Context, rdb *redis.Client) (*storage, error) {
	cmd := rdb.Ping(ctx)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", cmd.Err()) // TODO: handle error correctly
	}

	return &storage{
		rdb: rdb,
	}, nil
}

func (s *storage) Put(ctx context.Context, key string, page model.Page) error {
	panic("unimplemented")
}

func (s *storage) Get(ctx context.Context, key string) (model.Page, error) {
	panic("unimplemented")
}

func (s *storage) Delete(ctx context.Context, key string) error {
	panic("unimplemented")
}

func (s *storage) Top(ctx context.Context) ([]model.Page, error) {
	panic("unimplemented")
}
