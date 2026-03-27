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
	switch err {
	case nil:
		page = p
	case ErrNotFound:
	// Do nothing.
	default:
		return fmt.Errorf("failed to get data from the storage: %w", err)
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
	if err := s.storage.Delete(ctx, key); err != nil {
		return fmt.Errorf("failed to delete data from the storage: %w", err)
	}

	return nil
}

func (s *service) Top(ctx context.Context) ([]model.Page, error) {
	pages, err := s.storage.Top(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top pages from the storage: %w", err)
	}

	// TODO: I think a sorting mechanism has to be implemented there.

	return pages, nil
}
