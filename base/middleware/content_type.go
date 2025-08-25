package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONContentTypeMiddleware ensures all responses are JSON
func JSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

// ContentTypeJSON ensures Content-Type is application/json for requests with body
func ContentOnlyJSONMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Skip for GET, HEAD, OPTIONS and DELETE requests
		if ctx.Request.Method == http.MethodGet ||
			ctx.Request.Method == http.MethodHead ||
			ctx.Request.Method == http.MethodOptions ||
			ctx.Request.Method == http.MethodDelete {
			ctx.Next()
			return
		}

		contentType := ctx.Request.Header.Get("Content-Type")

		fmt.Println("Content Length ", ctx.Request.ContentLength)

		// Allow empty content type for requests without body
		// if ctx.Request.ContentLength == 0 && contentType == "" {
		// 	ctx.Next()
		// 	return
		// }

		if contentType != "application/json" && ctx.Request.ContentLength > 0 {
			ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Content-Type must be application/json",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
