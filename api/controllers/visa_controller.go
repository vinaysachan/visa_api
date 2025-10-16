package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/api/actions/implement"
	"github.com/vinaysachan/visa_api/api/request"
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

func (U *VisaController) ApplicationFormSubmit(context *gin.Context) {
	// Validate request
	request, ok := request.ValidateApplicationRequestData(context)
	if !ok {
		return
	}
	//Perform Bussiness Login :-
	resp, err := visaAppAction.VisaApplySave(*request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (U *VisaController) BasicDetailFormData(context *gin.Context) {
	//Validate Request

}
