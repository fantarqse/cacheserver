package storage

import (
	"context"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

type storage struct {
}

func New() *storage {
	return &storage{}
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

func (s *storage) Top(ctx context.Context) ([]string, error) {
	panic("unimplemented")
}
