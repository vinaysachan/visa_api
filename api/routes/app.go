package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/controllers"
	"github.com/vinaysachan/visa_api/api/middleware"
)

func SetupAppRoutes(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	v1.Use(middleware.ValidateDeviceIdMiddleware())
	{
		appController := controllers.NewAppController()

		v1.GET("/generate/csrf", appController.GenerateCsrfToken)

	}

}
