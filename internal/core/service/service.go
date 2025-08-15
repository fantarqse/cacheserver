package service

import (
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

func (s *service) Put(key string, page []byte) error {
	return nil
}

func (s *service) Get(key string) (model.Page, error) {
	page, err := s.storage.Get(key)
	if err != nil {
		return model.Page{}, err
	}

	return page, nil
}

func (s *service) Delete(key string) error {
	return s.storage.Delete(key)
}

func (s *service) Top() ([]string, error) {
	return nil, nil
}
