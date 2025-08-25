package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Simple salt – in production, use a per-user unique salt stored in DB
const globalSalt = "change_this_salt_123"

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(globalSalt + password))
	return hex.EncodeToString(hash.Sum(nil))
}

func ComparePassword(hashed, plain string) bool {
	return hashed == HashPassword(plain)
}
