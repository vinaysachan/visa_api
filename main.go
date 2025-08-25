package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/routes"
	"github.com/vinaysachan/visa_api/base/config"
	"github.com/vinaysachan/visa_api/base/middleware"
	"github.com/vinaysachan/visa_api/base/packages/passport"
	"github.com/vinaysachan/visa_api/base/utils"
)

// @title           Visa Rest API
// @version         1.0
// @description     A Rest Service for Visa
// @contact.name   API Support
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5005
// @BasePath  /api/v1
func main() {

	// Create Gin router
	router := gin.Default()

	//Set Logger & Recovery :-
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.RecoveryMiddleware())

	// Connect to main app database (run at startup)
	config.InitMainDB()

	// Setup routes
	routes.SetupAPIRoutes(router)

	//Setup Mode Application
	switch env := utils.GodotEnv("GO_ENV"); env {
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Route not found",
			"path":    c.Request.URL.Path,
			"method":  c.Request.Method,
			"message": "The requested resource could not be found",
		})
	})

	server := &http.Server{
		Addr:    ":" + utils.GodotEnv("GO_PORT"),
		Handler: router,
		// ReadTimeout:    25 * time.Second,
		// WriteTimeout:   30 * time.Second,
		// IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Create default personal access client if not exists
	passport.CreatePersonalAccessClient()

	//Run Server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
