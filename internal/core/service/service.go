package service

import (
	"context"

	"github.com/fantarqse/cacheserver/internal/core/model"
	"github.com/fantarqse/cacheserver/internal/core/port"
)

type service struct {
	storage port.CacheStorage
}

func New(storage port.CacheStorage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) Put(ctx context.Context, key string, page model.Page) error {
	if err := s.storage.Put(ctx, key, page); err != nil {
		return err
	}

	return nil
}

func (s *service) Get(ctx context.Context, key string) (model.Page, error) {
	page, err := s.storage.Get(ctx, key)
	if err != nil {
		return model.Page{}, err
	}

	return page, nil
}

func (s *service) Delete(ctx context.Context, key string) error {
	return s.storage.Delete(ctx, key)
}

func (s *service) Top(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (s *service) increaseHitRating(key string) (int, error) {
	return 0, nil
}
