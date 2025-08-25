package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vinaysachan/visa_api/base/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64               `gorm:"primaryKey;autoIncrement"`
	UUID      string               `json:"uuid" gorm:"primaryKey;unique;type:uuid;"`
	Email     string               `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password  string               `gorm:"type:varchar(100);not null"`
	Mobile    utils.StringOrNumber `gorm:"type:varchar(15);not null;uniqueIndex"`
	Name      string               `gorm:"type:varchar(100);not null"`
	Status    string               `gorm:"type:enum('Y','N');default:'Y';not null"`
	CreatedAt time.Time            `gorm:"not null"`
	UpdatedAt *time.Time
}

// TableName overrides the default table name
func (User) TableName() string {
	return "app_users"
}

// BeforeCreate generates UUID automatically before insert
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == "" {
		u.UUID = uuid.NewString()
	}
	return
}
