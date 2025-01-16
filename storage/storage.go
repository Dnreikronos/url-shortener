package storage

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
