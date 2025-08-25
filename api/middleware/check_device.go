package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateDeviceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceID := c.GetHeader("X-Device-ID") // client must send device ID
		if deviceID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing device Id headers"})
			return
		}
		deviceUUID, err := uuid.Parse(deviceID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Device-ID must be a valid UUID"})
			return
		}
		fmt.Println("deviceID", deviceID, "deviceUUID", deviceUUID)
		c.Set("X-Device-ID", deviceUUID)
		c.Next()
	}
}
