package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/actions/implement"
)

type VisaController struct{}

func NewVisaController() *VisaController {
	return &VisaController{}
}

var visaAppAction = implement.NewVisaAction()

func (U *VisaController) ApplicationFormData(context *gin.Context) {
	resp, err := visaAppAction.VisaApplyFormData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}
