package request

import (
	"github.com/gin-gonic/gin"
)

// Login request structure
type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=5,max=25,password_strength"`
}

func ValidateLoginUser(c *gin.Context) (*LoginUserRequest, bool) {
	var request LoginUserRequest
	if !GenericValidate(c, &request) {
		return nil, false
	}
	return &request, true
}

// Login response structure
type LoginResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"` // Seconds until expiration
	ExpiresAt    string `json:"expires_at"` // Formatted datetime string
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// CSRF Token request structure
// type CsrfTokenRequest struct {
// 	DeviceId string `json:"device_id" validate:"required,uuid,max=255"`
// }

// func ValidateCsrfRequest(c *gin.Context) (*CsrfTokenRequest, bool) {
// 	var request CsrfTokenRequest
// 	if !GenericValidate(c, &request) {
// 		return nil, false
// 	}
// 	return &request, true
// }

// CSRF Token response structure
type CsrfTokenResponse struct {
	ExpiresAt string `json:"expires_at"` // Formatted datetime string
	Token     string `json:"access_token"`
}
