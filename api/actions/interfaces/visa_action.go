package interfaces

import (
	"github.com/vinaysachan/visa_api/api/request"
)

type VisaAction interface {
	VisaApplyFormData() (*request.VisaApplyFormDataResponse, error)
	VisaApplySave(request.VisaApplicationDataRequest) (*request.VisaApplicationApplyResponse, error)
}
