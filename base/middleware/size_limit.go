package middleware

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

const (
	defaultMaxBodySize = 10 << 20 // 10 MB,  Default size limits,  Therefore, 10 << 20 is equivalent to 10 * (2^20).
	authMaxBodySize    = 1 << 20  // 1 MB, Default size limits for auth endpoints
	uploadMaxBodySize  = 50 << 20 // 50 MB, Default size limits for upload endpoints

)

func RequestPayloadSizeLimit() gin.HandlerFunc {
	// Default size limit for most endpoints
	return limits.RequestSizeLimiter(defaultMaxBodySize)
}

func AuthRequestPayloadSizeLimiter() gin.HandlerFunc {
	// Smaller size limit for auth endpoints
	return limits.RequestSizeLimiter(authMaxBodySize)
}

func UploadRequestPayloadSizeLimiter() gin.HandlerFunc {
	// Larger size limit for upload endpoints
	return limits.RequestSizeLimiter(uploadMaxBodySize)
}
