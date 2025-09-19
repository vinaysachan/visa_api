package request

import (
	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/base/utils"
)

type CountryResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ArrivalPortResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type VisaTypeResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ApplicationTypeResponse struct {
	ID        uint64             `json:"id"`
	Name      string             `json:"name"`
	VisaTypes []VisaTypeResponse `json:"visa_types"`
}

// Visa Apply Form Data Response structure
type VisaApplyFormDataResponse struct {
	Countries        []CountryResponse         `json:"countries"`
	ArrivalPorts     []ArrivalPortResponse     `json:"arrival_ports"`
	ApplicationTypes []ApplicationTypeResponse `json:"application_types"`
}

type VisaApplicationDataRequest struct {
	ApplicationType utils.StringOrNumber `json:"application_type" validate:"required"`
	Fname           string               `json:"fname" validate:"required,min=2,max=100"`
	Mname           *string              `json:"mname" validate:"omitempty,min=2,max=100"`
	Lname           string               `json:"lname" validate:"required,min=2,max=100"`
	PassportType    string               `json:"passport_type" validate:"required"`
	Nationality     uint64               `json:"nationality"  validate:"required"`
	PortOfArrival   uint64               `json:"portofarrival" validate:"required"`
	PassportNumber  utils.StringOrNumber `json:"passport_number" validate:"required"`
	VisaType        uint64               `json:"visa_type" validate:"required"`
	DateOfBirth     string               `json:"date_of_birth" validate:"required,datetime=2006-01-02,before_today,min_age=1"`
	ArrivalDate     string               `json:"arrival_date" validate:"required,datetime=2006-01-02,before_today"`
	Phone           utils.StringOrNumber `json:"phone" validate:"required,min=10,max=15,valid_mobile_number"`
	Email           string               `json:"email" validate:"required,email,max=255"`
}

func ValidateApplicationRequestData(c *gin.Context) (*VisaApplicationDataRequest, bool) {
	var request VisaApplicationDataRequest
	if !GenericValidate(c, &request) {
		return nil, false
	}
	return &request, true
}

type VisaApplicationApplyResponse struct {
	Message       string `json:"message"`
	ID            uint64 `json:"id"`
	ApplicationId string `json:"application_id"`
}
