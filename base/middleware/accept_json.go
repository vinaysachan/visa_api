package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AcceptJSONMiddleware ensures the client accepts JSON responses
func AcceptJSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptHeader := c.GetHeader("Accept")

		// Skip for OPTIONS requests (preflight)
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// Check if client accepts JSON
		if acceptHeader != "application/json" {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": "This API only serves application/json responses",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
