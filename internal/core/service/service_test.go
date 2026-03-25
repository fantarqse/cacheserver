package service

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

type mockStorage struct {
	data map[string]model.Page
}

func (m *mockStorage) Put(_ context.Context, key string, page model.Page) error {
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

func (m *mockStorage) Top(_ context.Context) ([]string, error) {
	return nil, nil
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
			m := &mockStorage{
				data: tt.state,
			}
			svc := New(m)

			got, err := svc.Get(context.Background(), tt.key)

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
		})
	}
}
