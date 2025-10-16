package storage

import "github.com/fantarqse/cacheserver/internal/core/model"

type storage struct {
}

func New() *storage {
	return &storage{}
}

func (s *storage) Put(key string, page model.Page) error {
	panic("unimplemented")
}

func (s *storage) Get(key string) (model.Page, error) {
	panic("unimplemented")
}

func (s *storage) Delete(key string) error {
	panic("unimplemented")
}

func (s *storage) Top() ([]string, error) {
	panic("unimplemented")
}
