package storage

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage interface {
	SaveURL(shortKey, OiriginalKey string)error
	GetOriginalURL(shortKey string) (string, error)
}

type GormStorage struct {
	db *gorm.DB
}

type URLMapping struct {
	id          uint   `gorm:"primaryKey"`
	ShortKey    string `gorm:"uniqueIndex"`
	OriginalURL string
}

var (
	db    *gorm.DB
	mutex = &sync.Mutex{}
)

func InitializeDatabase() {
	var err error
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&URLMapping{})
}

func SaveURL(shortKey, originalURL string) {
	mutex.Lock()
	defer mutex.Unlock()

	urlMapping := URLMapping{ShortKey: shortKey, OriginalURL: originalURL}
	db.Create(&urlMapping)
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
