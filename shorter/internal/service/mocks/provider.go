package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ProviderMock struct {
	mock.Mock
}

func (s *ProviderMock) Create(ctx context.Context, url string) (string, error) {
	args := s.Called(ctx, url)

	return args.String(0), args.Error(1)
}

func (s *ProviderMock) GetByID(ctx context.Context, id string) (string, error) {
	args := s.Called(ctx, id)

	return args.String(0), args.Error(1)
}

func (s *ProviderMock) Close() {
}
