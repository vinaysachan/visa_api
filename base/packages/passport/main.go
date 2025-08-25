package passport

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/vinaysachan/visa_api/base/config"
	"github.com/vinaysachan/visa_api/base/utils"
	"gorm.io/gorm"
)

func CreatePersonalAccessClient() {
	var client OAuthClient
	db := config.MainDB

	result := db.First(&client, utils.GodotEnv("OAUTH_CLIENT_ID"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		client = OAuthClient{
			ID:                   utils.GodotEnv("OAUTH_CLIENT_ID"),
			Name:                 "Personal Access Client",
			Secret:               utils.GodotEnv("OAUTH_CLIENT_SECRET"),
			Redirect:             "http://localhost",
			PersonalAccessClient: true,
			PasswordClient:       false,
			Revoked:              false,
		}
		db.Create(&client)
	}
}

func GenerateToken() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GetTokenData retrieves and validates token, returning the token data
func GetTokenData(token string) (*OAuthAccessToken, error) {
	// Expecting "Bearer <token>"
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid authorization passport header format")
	}

	accessToken := parts[1]

	// Find the token in database
	var authToken OAuthAccessToken

	result := config.MainDB.Unscoped().Where("`id` = ?", accessToken).First(&authToken)
	if result.Error != nil {
		return nil, errors.New("invalid access token")
	}

	if authToken.Revoked || (authToken.ExpiresAt != nil && time.Now().After(*authToken.ExpiresAt)) {
		return nil, errors.New("access token revoked or expired")
	}

	return &authToken, nil
}

func SaveAccessTokenInDB(authToken *OAuthAccessToken) error {
	//Before save token, check client exists
	var client OAuthClient
	result := config.MainDB.Unscoped().Where("`id` = ?", authToken.ClientID).First(&client)
	if result.Error != nil {
		return errors.New("invalid client")
	}
	if config.MainDB.Create(authToken).Error != nil {
		return errors.New("could not generate token")
	}
	return nil
}

func SaveRefreshTokenInDB(refreshToken *OAuthRefreshToken) error {
	if config.MainDB.Create(refreshToken).Error != nil {
		return errors.New("could not generate refresh token")
	}
	return nil
}

func RevokeTokenInDB(accessTokenID string) error {
	var authToken OAuthAccessToken
	return config.MainDB.Model(&authToken).Where("`id` = ?", accessTokenID).Update("revoked", true).Error
}

func RevokeUserTokenInDB(UserID uint64, accessTokenID string) error {
	var authToken OAuthAccessToken
	return config.MainDB.Model(&authToken).Where("`id` != ?", accessTokenID).Where("`user_id` = ?", UserID).Update("revoked", true).Error
}

func SaveCsrfTokenInDB(csrfToken *OAuthCsrfClaim) error {
	if config.MainDB.Create(csrfToken).Error != nil {
		return errors.New("could not generate csrf token")
	}
	return nil
}
