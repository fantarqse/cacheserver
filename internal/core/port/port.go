package port

import (
	"context"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

// Port for the drivers
type CacheService interface {
	Put(ctx context.Context, key string, page model.Page) error
	Get(ctx context.Context, key string) (model.Page, error)
	Delete(ctx context.Context, key string) error
	Top(ctx context.Context) ([]model.Page, error)
}

// Port for the driven
type CacheStorage interface {
	Put(ctx context.Context, key string, page model.Page) error
	Get(ctx context.Context, key string) (model.Page, error)
	Delete(ctx context.Context, key string) error
	Top(ctx context.Context) ([]model.Page, error)
}
