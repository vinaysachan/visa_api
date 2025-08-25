package config

import (
	"log"

	"github.com/vinaysachan/visa_api/base/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MainDB *gorm.DB

// Connect to main app database (run at startup)
func InitMainDB() {
	dsn := utils.GodotEnv("DATABASE_URI_PROD")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to main DB: %v", err)
	}
	MainDB = db
}
