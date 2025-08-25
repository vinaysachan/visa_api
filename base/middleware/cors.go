package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/base/utils"
)

// Define constants for environment names to improve readability and prevent typos.
const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvLocal       = "local"
	EnvProduction  = "production" // Explicitly define production for clarity
)

// getAllowedOriginsForEnvironment determines the allowed CORS origins based on the current environment.
// This function adheres to the Single Responsibility Principle.
func getAllowedOriginsForEnvironment(env string) []string {
	switch env {
	case EnvDevelopment, EnvTest, EnvLocal:
		return []string{
			"http://localhost:3000", // React default
			"http://localhost:8080", // Vue default
			// Add other local development origins as needed
		}
	case EnvProduction:
		return []string{
			"https://your-production-domain.com",
			"https://www.your-production-domain.com",
			// Add other production domains
		}
	default:
		// Default to a restrictive policy or production if environment is unknown
		// Or, you might want to return an empty slice or log an error for unknown environments
		return []string{} // Most restrictive by default for unknown environments
	}
}

// CORSMiddleware provides a robust CORS configuration for the Gin framework.
// It leverages environment variables to dynamically set allowed origins.
func CORSMiddleware() gin.HandlerFunc {
	// Retrieve the current environment from your utility function.
	currentEnv := utils.GodotEnv("GO_ENV")

	// Initialize the base CORS configuration.
	// These settings are generally static across environments.
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,          // Set to true only if you need to handle cookies/credentials
		MaxAge:           12 * time.Hour, // Cache preflight requests for 12 hours
	}

	// Dynamically set allowed origins based on the environment.
	config.AllowOrigins = getAllowedOriginsForEnvironment(currentEnv)

	// If no specific origins are set for an environment (e.g., unknown env),
	// you might want to explicitly disallow all or log a warning.
	if len(config.AllowOrigins) == 0 {
		// Log a warning or error if no origins are configured for the current environment
		// fmt.Printf("WARNING: No CORS origins configured for environment: %s\n", currentEnv)
		// As a fallback, you might consider allowing all for development or none for production
		// For production, it's safer to explicitly list allowed origins.
	}

	return cors.New(config)
}
