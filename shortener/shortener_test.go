package shortener

import (
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) SaveURL(shortKey, originalURL string) error {
	args := m.Called(shortKey, originalURL)
	return args.Error(0)
}

func (m *MockStorage) GetOriginalURL(shortKey string) (string, error) {
	args := m.Called(shortKey)
	return args.String(0), args.Error(1)
}
