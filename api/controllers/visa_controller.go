package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VisaController struct{}

func NewVisaController() *VisaController {
	return &VisaController{}
}

// var authAppAction = implement.NewAuthAction()

func (U *VisaController) ApplicationFormData(context *gin.Context) {

	context.JSON(http.StatusOK, nil)
}
