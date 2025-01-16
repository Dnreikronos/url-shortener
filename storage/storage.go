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
