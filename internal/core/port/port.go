package port

import "github.com/fantarqse/cacheserver/internal/core/model"

// Port for the drivers
type CacheService interface {
	Put(key string, page []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Top() ([]string, error)
}

// Port for the driven
type CacheStorage interface {
	Put(key string, page model.Page) error
	Get(key string) (model.Page, error)
	Delete(key string) error
	Top() ([]string, error)
}

