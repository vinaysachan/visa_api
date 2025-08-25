package interfaces

import (
	"github.com/vinaysachan/visa_api/api/request"
)

type VisaAction interface {
	VisaApplyFormData() (*request.VisaApplyFormDataResponse, error)
}
