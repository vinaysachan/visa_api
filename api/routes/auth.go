package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/controllers"
	"github.com/vinaysachan/visa_api/api/middleware"
)

func SetupUsersRoutes(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	{
		userController := controllers.NewUserController()

		v1.POST("/user/login", userController.LoginWithPassword)
		v1.POST("/user/logout", middleware.AuthMiddleware(), userController.LogoutUser)

		//
		// v1.POST("/user/resister", userController.RegisterNewUser)
		//
		// v1.POST("/user/logout", middleware.AuthMiddleware(), userController.LogoutUser)
		// v1.POST("/user/resister", userController.VerifyUser)
		// v1.POST("/user/resister", userController.UpdateUser)
		// v1.POST("/user/resister", userController.DeleteUser)
		// v1.POST("/user/resister", userController.UpdateUserPassword)
		// v1.POST("/user/resister", userController.UpdateUserEmail)
		// v1.POST("/user/resister", userController.UpdateUserPhone)
		// v1.POST("/user/resister", userController.UpdateUserAddress)
		// v1.POST("/user/resister", userController.UpdateUserBank)
		// v1.POST("/user/resister", userController.UpdateUserBankAccount)
	}

}
