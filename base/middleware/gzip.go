package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func GzipMiddleware() gin.HandlerFunc {
	return gzip.Gzip(
		gzip.BestCompression, //level 9, This provides the highest compression ratio at the cost of slightly more CPU usage
		//gzip.DefaultCompression, // level 6, This provides the lowest compression ratio at the cost of slightly less CPU usage
		//gzip.BestSpeed, // level 1, This provides the lowest compression ratio at the cost of slightly less CPU usage

		// Add more excluded paths
		gzip.WithExcludedPaths([]string{"/metrics", "/health"}),

		// Add more excluded extensions
		gzip.WithExcludedExtensions([]string{".png", ".gif", ".jpeg", ".jpg", ".pdf"}),
	)
}
