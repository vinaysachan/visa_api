package passport

import (
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost int64 = 12

// Helper to generate random tokens
// func generateToken() (string, error) {
// 	b := make([]byte, 32)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(b), nil
// }

// Hash password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), int(bcryptCost))
	return string(bytes), err
}

// Check password hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
