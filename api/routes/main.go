package routes

import (
	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/base/middleware"
)

func SetupAPIRoutes(router *gin.Engine) *gin.Engine {

	// Use default security headers
	router.Use(ginhelmet.Default())

	// Apply Gzip compression middleware first
	router.Use(middleware.GzipMiddleware())

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Global middleware
	router.Use(middleware.AcceptJSONMiddleware()) //AcceptJSONMiddleware ensures the client accepts JSON responses

	// start background cleanup
	middleware.InitThrottleCleanup()

	api := router.Group("/api")
	api.Use(middleware.ContentOnlyJSONMiddleware())
	{
		authAPIRoutes := api.Group("")
		authAPIRoutes.Use(middleware.AuthRequestPayloadSizeLimiter())
		authAPIRoutes.Use(middleware.Throttle("10,1")) // 10 req/minute
		{
			SetupUsersRoutes(authAPIRoutes)
		}

		apiRoutes := api.Group("")
		apiRoutes.Use(middleware.RequestPayloadSizeLimit()) // Apply default size limit middleware
		apiRoutes.Use(middleware.Throttle("20,1"))          // 20 req/minute
		{
			SetupAppRoutes(apiRoutes)
			SetupVisaRoutes(apiRoutes)
		}

	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
