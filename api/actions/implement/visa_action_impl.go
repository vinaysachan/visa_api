package implement

import (
	"github.com/vinaysachan/visa_api/api/actions/interfaces"
	"github.com/vinaysachan/visa_api/api/request"
)

type visaActionImpl struct{}

func NewVisaAction() interfaces.VisaAction {
	return &visaActionImpl{}
}

func (v *visaActionImpl) VisaApplyFormData() (*request.VisaApplyFormDataResponse, error) {
	// Initialize with empty slices for countries and ports
	return &request.VisaApplyFormDataResponse{
		Countries: []request.CountryResponse{},
		Ports:     []request.PortResponse{},
	}, nil
}
