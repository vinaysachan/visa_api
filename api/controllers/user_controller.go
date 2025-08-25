package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/actions/implement"
	"github.com/vinaysachan/visa_api/api/request"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var authAction = implement.NewAuthAction()

func (U *UserController) LoginWithPassword(context *gin.Context) {
	// Validate request
	request, ok := request.ValidateLoginUser(context)
	if !ok {
		return
	}

	resp, err := authAction.LoginUser(*request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, resp)
}

func (U *UserController) LogoutUser(context *gin.Context) {
	accessTokenID, exists := context.Get("accessToken")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Revoke the token
	resp, err := authAction.RevokeToken(accessTokenID.(string))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, resp)
}
