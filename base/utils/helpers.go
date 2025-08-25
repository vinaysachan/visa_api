package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// Helper function to generate random tokens
func GenerateToken() (string, error) {
	b := make([]byte, 40)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
