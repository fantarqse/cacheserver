package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/fantarqse/cacheserver/internal/core/model"
	"github.com/fantarqse/cacheserver/internal/core/port"
)

var (
	ErrNotFound = errors.New("key doesn't exist")
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
	p, err := s.storage.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return fmt.Errorf("failed to get data from the storage: %w", err)
		}
		page = p
	}
	// * If a page already exists in the storage,
	//   it rebinds, and the hit rating increases.
	// * If it doesn't exist, the hit rating just becomes 1.
	page.HitRating++

	if err := s.storage.Put(ctx, key, page); err != nil {
		return fmt.Errorf("failed to put data to the storage: %w", err)
	}

	return nil
}

func (s *service) Get(ctx context.Context, key string) (model.Page, error) {
	page, err := s.storage.Get(ctx, key)
	if err != nil {
		return model.Page{}, err
	}

	page.HitRating++

	// TODO: could I sync data concurrently there?
	// I don't want to block the program execution just for updating.
	if err := s.storage.Put(ctx, key, page); err != nil {
		return model.Page{}, err
	}

	return page, nil
}

func (s *service) Delete(ctx context.Context, key string) error {
	return s.storage.Delete(ctx, key)
}

func (s *service) Top(ctx context.Context) ([]string, error) {
	// TODO: I think a sorting mechanism has to be implemented there.
	return nil, nil
}

func (s *service) increaseHitRating(key string) (int, error) {
	return 0, nil
}
