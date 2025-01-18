package storage

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage interface {
	SaveURL(shortKey, OiriginalKey string) error
	GetOriginalURL(shortKey string) (string, error)
}

type GormStorage struct {
	db *gorm.DB
}

type URLMapping struct {
	ID          uint   `gorm:"primaryKey"`
	ShortKey    string `gorm:"uniqueIndex"`
	OriginalURL string
}

var (
	db    *gorm.DB
	mutex = &sync.Mutex{}
)

func NewGormStorage() (*GormStorage, error) {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&URLMapping{})
	return &GormStorage{db: db}, nil
}

func (s *GormStorage) SaveURL(shortKey, originalURL string) error {
	urlMapping := URLMapping{ShortKey: shortKey, OriginalURL: originalURL}
	result := s.db.Create(&urlMapping)
	return result.Error
}

func GetOriginalURL(shortKey string) string {
	mutex.Lock()
	defer mutex.Unlock()

	var urlMapping URLMapping
	result := db.First(&urlMapping, "short_key = ?", shortKey)
	if result.Error != nil {
		return ""
	}
	return urlMapping.OriginalURL
}
