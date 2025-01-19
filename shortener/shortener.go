package shortener

import (
	"fmt"
	"hash/fnv"

	"github.com/Dnreikronos/url-shortener/storage"
)

type URLShortener struct {
	storage storage.Storage
}

func NewURLShortener(storage storage.Storage) *URLShortener {
	return &URLShortener{storage: storage}
}

func (us *URLShortener) ShortenURL(url string) (string, error) {
	hash := hashURL(url)
	shortKey := fmt.Sprintf("%d", hash)

	err := us.storage.SaveURL(shortKey, url)
	if err != nil {
		return "", err
	}
	return shortKey, nil
}

func ExpandURL(shortenedURL string) string {
	return storage.GetOriginalURL(shortenedURL)
}

func hashURL(url string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(url))
	return h.Sum32()
}
