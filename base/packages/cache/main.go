package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/vinaysachan/visa_api/base/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Set stores a value with TTL
func Set(key string, value interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cacheModal := CacheEntry{
		CacheKey:  key,
		Data:      jsonData,
		ExpiresAt: time.Now().Add(ttl),
	}

	upsertData := config.MainDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "cache_key"}},                                   // unique column
		DoUpdates: clause.AssignmentColumns([]string{"data", "expires_at", "updated_at"}), // fields to update if conflict
	}).Create(&cacheModal).Error

	if upsertData != nil {
		return errors.New("could not save cache")
	}

	return upsertData
}

// Get retrieves cached data into dest
func Get(key string, dest interface{}) error {
	var cacheModal CacheEntry
	err := config.MainDB.Where("cache_key = ?", key).First(&cacheModal).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("cache: key not found")
		}
		return err
	}

	if time.Now().After(cacheModal.ExpiresAt) {
		// delete expired
		config.MainDB.Delete(&cacheModal)
		return errors.New("cache: key expired")
	}
	return json.Unmarshal(cacheModal.Data, dest)
}

// Delete removes a cache entry
func Delete(key string) error {
	err := config.MainDB.Where("cache_key = ?", key).Delete(&CacheEntry{}).Error
	return err
}

// Cleanup expired cache
func Cleanup() error {
	return config.MainDB.Where("expires_at < ?", time.Now()).Delete(&CacheEntry{}).Error
}

// Remember checks cache, if not found executes callback, stores result and returns it.
func Remember[T any](key string, ttl time.Duration, callback func() (T, error)) (T, error) {
	var result T

	// Try to get from cache
	err := Get(key, &result)
	if err == nil {
		return result, nil
	}

	// Cache miss or expired → call callback
	value, err := callback()
	if err != nil {
		var zero T
		return zero, err
	}

	// Save in cache
	_ = Set(key, value, ttl)

	return value, nil
}
