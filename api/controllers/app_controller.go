package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinaysachan/visa_api/api/actions/implement"
)

type AppController struct{}

func NewAppController() *AppController {
	return &AppController{}
}

var authAppAction = implement.NewAuthAction()

func (U *AppController) GenerateCsrfToken(context *gin.Context) {

	deviceID, exist := context.Get("X-Device-ID")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Device ID is missing"})
		return
	}

	resp, err := authAppAction.CsrfToken(deviceID.(uuid.UUID), context.ClientIP())
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusCreated, resp)
}
