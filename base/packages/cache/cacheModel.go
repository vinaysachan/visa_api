package cache

import (
	"time"
)

type CacheEntry struct {
	CacheKey  string    `gorm:"uniqueIndex;size:255;not null"`
	Data      []byte    `gorm:"type:json;not null"`
	ExpiresAt time.Time `gorm:"index;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CacheEntry) TableName() string {
	return "cache"
}
