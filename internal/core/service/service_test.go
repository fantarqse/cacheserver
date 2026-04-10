package service

import (
	"context"
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/fantarqse/cacheserver/internal/core/model"
	"github.com/fantarqse/cacheserver/internal/platform/flags"
)

type mockStorage struct {
	data map[string]model.Page
}

func (m *mockStorage) Put(_ context.Context, key string, page model.Page, ttl time.Duration) error {
	m.data[key] = page
	return nil
}

func (m *mockStorage) Get(_ context.Context, key string) (model.Page, error) {
	p, ok := m.data[key]
	if !ok {
		return model.Page{}, ErrNotFound
	}

	return p, nil
}

func (m *mockStorage) Delete(_ context.Context, key string) error {
	delete(m.data, key)
	return nil
}

func (m *mockStorage) Top(_ context.Context) ([]model.Page, error) {
	return nil, nil
}

func TestPut(t *testing.T) {
	tests := []struct {
		name           string
		state          map[string]model.Page
		key            string
		data           model.Page
		expectedData   []byte
		expectedRating int
		expectedErr    error
		beforeErr      error
	}{
		{
			name: "1. Successful: adding data to the empty storage",
			state: func() map[string]model.Page {
				m := make(map[string]model.Page)
				return m
			}(),
			key:            "key1",
			data:           model.Page{Data: []byte("data1")},
			expectedData:   []byte("data1"),
			expectedRating: 1,
			expectedErr:    nil,
			beforeErr:      ErrNotFound,
		},
		{
			name: "2. Successful: updating data in the storage",
			state: func() map[string]model.Page {
				m := make(map[string]model.Page)
				m["key1"] = model.Page{
					Data:      []byte("data1"),
					HitRating: 2,
				}
				return m
			}(),
			key:            "key1",
			data:           model.Page{Data: []byte("data1")},
			expectedData:   []byte("data1"),
			expectedRating: 3,
			expectedErr:    nil,
			beforeErr:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err error
				cfg flags.Cache

				ctx = context.Background()
				m   = &mockStorage{data: tt.state}
				svc = New(cfg, m)
			)

			// Before
			_, err = m.Get(ctx, tt.key)
			if !errors.Is(tt.beforeErr, err) {
				t.Fatalf("before operation: expected %q, but got: %q", tt.beforeErr, err)
			}

			// During
			err = svc.Put(ctx, tt.key, tt.data)
			if tt.expectedErr == nil && err != nil {
				t.Fatalf("expected no error, but got: %q", err)
			}
			if tt.expectedErr != nil && err == nil {
				t.Fatalf("expected %q, but got nil", tt.expectedErr)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected %q, but got %q", tt.expectedErr, err)
			}

			// After
			after, err := m.Get(ctx, tt.key)
			if err != nil {
				t.Fatalf("after: expected no error, but got %q", err)
			}

			if tt.expectedRating != after.HitRating {
				t.Fatalf("after: expected rating %q, but got %q", tt.expectedRating, after.HitRating)
			}
			if !slices.Equal(tt.expectedData, after.Data) {
				t.Fatalf("after: expected data %q, but got %q", tt.expectedData, after.Data)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		state          map[string]model.Page
		key            string
		expectedData   []byte
		expectedRating int
		expectedErr    error
	}{
		{
			name: "1. Successful",
			state: func() map[string]model.Page {
				m := make(map[string]model.Page)

				m["key1"] = model.Page{
					Data:      []byte("data1"),
					HitRating: 1,
				}
				m["key2"] = model.Page{
					Data:      []byte("data2"),
					HitRating: 10,
				}

				return m
			}(),
			key:            "key1",
			expectedData:   []byte("data1"),
			expectedRating: 2,
			expectedErr:    nil,
		},
		{
			name:           "2. Failed: a key doesn't exist in the storage",
			state:          make(map[string]model.Page),
			key:            "key1",
			expectedData:   nil,
			expectedRating: 0,
			expectedErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err error
				cfg flags.Cache

				ctx = context.Background()
				m   = &mockStorage{data: tt.state}
				svc = New(cfg, m)
			)

			got, err := svc.Get(ctx, tt.key)

			// Errors
			if tt.expectedErr == nil && err != nil {
				t.Fatalf("expected no error, but got: %q", err)
			}
			if tt.expectedErr != nil && err == nil {
				t.Fatalf("expected %q, but got nil", tt.expectedErr)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected %q, but got %q", tt.expectedErr, err)
			}

			// Data
			if !slices.Equal(tt.expectedData, got.Data) {
				t.Fatalf("expected %q, but got %q", tt.expectedData, got.Data)
			}
			if tt.expectedRating != got.HitRating {
				t.Fatalf("expected %v, but got %v", tt.expectedRating, got.HitRating)
			}

			// Checking if Get method updates value in the storage.
			synced, err := m.Get(ctx, tt.key)
			if tt.expectedErr == nil && err != nil {
				t.Fatalf("expected no err after sync, but got: %v", err)
			}

			if !slices.Equal(synced.Data, got.Data) {
				t.Fatal("data must be equal after sync")
			}
			if synced.HitRating != got.HitRating {
				t.Fatal("hit rating must be equal after sync")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name        string
		state       map[string]model.Page
		key         string
		expectedErr error
		beforeErr   error
	}{
		{
			name: "1. Successful: a key is in the storage",
			state: func() map[string]model.Page {
				m := make(map[string]model.Page)
				m["key1"] = model.Page{}
				return m
			}(),
			key:         "key1",
			expectedErr: nil,
			beforeErr:   nil,
		},
		{
			name: "2. Successful: a key is not in the storage",
			state: func() map[string]model.Page {
				m := make(map[string]model.Page)
				return m
			}(),
			key:         "key1",
			expectedErr: nil,
			beforeErr:   ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err error
				cfg flags.Cache

				ctx = context.Background()
				m   = &mockStorage{data: tt.state}
				svc = New(cfg, m)
			)

			// Before
			_, err = m.Get(ctx, tt.key)
			if !errors.Is(tt.beforeErr, err) {
				t.Fatalf("before operation: expected %q, but got: %q", tt.beforeErr, err)
			}

			// During
			err = svc.Delete(ctx, tt.key)
			if tt.expectedErr != nil && err == nil {
				t.Fatalf("during: expected %q, but got nil", tt.expectedErr)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("during: expected %q, but got %q", tt.expectedErr, err)
			}

			// After
			_, err = m.Get(ctx, tt.key)
			if !errors.Is(err, ErrNotFound) {
				t.Fatalf("after: error must be %q, but got %q", ErrNotFound, err)
			}
		})
	}
}
