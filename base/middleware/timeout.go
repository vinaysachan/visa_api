package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace request with context
		c.Request = c.Request.WithContext(ctx)

		// Channel to track request completion
		done := make(chan struct{})
		defer close(done)

		// Track panic
		var err error
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()

		// Process request in goroutine
		go func() {
			defer func() {
				if r := recover(); r != nil {
					err = r.(error)
				}
			}()
			c.Next()
			done <- struct{}{}
		}()

		// Wait for completion or timeout
		select {
		case <-done:
			return
		case <-ctx.Done():
			if err == nil {
				err = ctx.Err()
			}
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"error": "Request timeout",
			})
			return
		}
	}
}
