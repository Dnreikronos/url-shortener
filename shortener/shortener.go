package shortener

import (
	"fmt"
	"hash/fnv"

	"github.com/Dnreikronos/url-shortener/storage"
)

func ShortenURL(url string) string {
type URLShortener struct {
	storage storage.Storage
}

func NewURLShortener(storage storage.Storage) *URLShortener {
	return &URLShortener{storage: storage}
}

	hash := hashURL(url)
	shortKey := fmt.Sprintf("%d", hash)

	storage.SaveURL(shortKey, url)
	return shortKey
}

func ExpandURL(shortenedURL string) string {
	return storage.GetOriginalURL(shortenedURL)
}

func hashURL(url string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(url))
	return h.Sum32()
}
