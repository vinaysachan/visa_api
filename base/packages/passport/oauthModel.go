package passport

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OAuthClient struct {
	ID                   string    `gorm:"primaryKey"`
	UserID               uint      `gorm:"index"`
	Name                 string    `gorm:"not null"`
	Secret               string    `gorm:"not null"`
	Redirect             string    `gorm:"type:text;not null"`
	PersonalAccessClient bool      `gorm:"not null"`
	PasswordClient       bool      `gorm:"not null"`
	Revoked              bool      `gorm:"not null"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
	Provider             *string
}

func (OAuthClient) TableName() string {
	return "oauth_clients"
}

type OAuthAccessToken struct {
	ID        string      `gorm:"primaryKey;size:100"`
	UserID    uint64      `gorm:"index"`
	ClientID  string      `gorm:"size:100;not null"`
	Scopes    *string     `gorm:"type:text"`
	Revoked   bool        `gorm:"not null"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
	Client    OAuthClient `gorm:"foreignKey:ClientID"`
	ExpiresAt *time.Time
	Name      *string
}

func (OAuthAccessToken) TableName() string {
	return "oauth_access_tokens"
}

type OAuthRefreshToken struct {
	ID            string           `gorm:"primaryKey;size:100"`
	AccessTokenID string           `gorm:"size:100;not null"`
	Revoked       bool             `gorm:"not null"`
	AccessToken   OAuthAccessToken `gorm:"foreignKey:AccessTokenID"`
	// CreatedAt     time.Time        `gorm:"autoCreateTime"`
	// UpdatedAt     time.Time        `gorm:"autoUpdateTime"`
	ExpiresAt *time.Time
}

func (OAuthRefreshToken) TableName() string {
	return "oauth_refresh_tokens"
}

type PersonalAccessClient struct {
	ID     uint
	Secret string
}

type OAuthCsrfClaim struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	DeviceID  uuid.UUID  `gorm:"type:char(36);not null;index" json:"device_id"`
	IP        string     `gorm:"type:varchar(500);not null" json:"ip"`
	Token     string     `gorm:"type:varchar(500);not null;uniqueIndex" json:"token"`
	ExpireAt  *time.Time `gorm:"type:timestamp" json:"expire_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (OAuthCsrfClaim) TableName() string {
	return "oauth_csrf_claims"
}

// BeforeCreate hook to auto-generate UUIDs and set default ExpireAt
func (d *OAuthCsrfClaim) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	if d.ExpireAt == nil {
		defaultExpire := time.Now().Add(24 * time.Hour) // Set default expiration to 24 hours from now
		d.ExpireAt = &defaultExpire
	}
	return
}
