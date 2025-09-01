package implement

import (
	"errors"

	"github.com/vinaysachan/visa_api/api/actions/interfaces"
	"github.com/vinaysachan/visa_api/api/request"
	"github.com/vinaysachan/visa_api/data/tasks"
)

type visaActionImpl struct{}

func NewVisaAction() interfaces.VisaAction {
	return &visaActionImpl{}
}

func (v *visaActionImpl) VisaApplyFormData() (*request.VisaApplyFormDataResponse, error) {

	//Get Active country list form task :-
	activeCountries, err := tasks.GetActiveCountryList()
	if err != nil {
		return nil, errors.New("failed to get active country list")
	}
	//Map activeCountries to the expected type
	countries := make([]request.CountryResponse, len(*activeCountries))
	for i, country := range *activeCountries {
		countries[i] = request.CountryResponse{ID: country.ID, Name: country.Name}
	}

	//Get Active Arrival port list form task :-
	activeArrivalPorts, err := tasks.GetActiveArrivalPortList()
	if err != nil {
		return nil, errors.New("failed to get active arrival port list")
	}

	//Map activeArrivalPorts to the expected type
	arrivalPorts := make([]request.ArrivalPortResponse, len(*activeArrivalPorts))
	for i, port := range *activeArrivalPorts {
		arrivalPorts[i] = request.ArrivalPortResponse{ID: port.ID, Name: port.Name}
	}

	//Get Application Type :-
	activeApplicationTypes, err := tasks.GetActiveVisaApplicationTypes()
	if err != nil {
		return nil, errors.New("failed to get visa application types")
	}

	//Map activeArrivalPorts to the expected type
	applicationTypes := make([]request.ApplicationTypeResponse, len(*activeApplicationTypes))
	for i, applicationType := range *activeApplicationTypes {
		visaTypes := make([]request.VisaTypeResponse, len(applicationType.VisaTypes))
		for j, visaType := range applicationType.VisaTypes {
			visaTypes[j] = request.VisaTypeResponse{ID: visaType.ID, Name: visaType.Name}
		}
		applicationTypes[i] = request.ApplicationTypeResponse{ID: applicationType.ID, Name: applicationType.Name, VisaTypes: visaTypes}
	}

	return &request.VisaApplyFormDataResponse{
		Countries:        countries,
		ArrivalPorts:     arrivalPorts,
		ApplicationTypes: applicationTypes,
	}, nil
}
