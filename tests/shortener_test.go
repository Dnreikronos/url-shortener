package tests

import (
	"testing"

	"github.com/Dnreikronos/url-shortener/shortener"
	"github.com/stretchr/testify/assert"
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

func TestShortenURL(t *testing.T) {
	mockStorage := new(MockStorage)
	shortener := shortener.NewURLShortener(mockStorage)

	mockStorage.On("SaveURL", mock.Anything, "https://example.com").Return(nil)

	shortKey, err := shortener.ShortenURL("https://example.com")

	assert.NoError(t, err)
	assert.NotEmpty(t, shortKey)
	mockStorage.AssertCalled(t, "SaveURL", mock.Anything, "https://example.com")
}
