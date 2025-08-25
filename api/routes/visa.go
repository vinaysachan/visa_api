package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/controllers"
	"github.com/vinaysachan/visa_api/api/middleware"
)

func SetupVisaRoutes(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	v1.Use(middleware.ValidateDeviceIdMiddleware())
	{
		visaController := controllers.NewVisaController()

		v1.GET("/evisa/form_data", visaController.ApplicationFormData)
	}

}
