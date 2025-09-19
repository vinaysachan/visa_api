package implement

import (
	"errors"
	"fmt"
	"time"

	"github.com/vinaysachan/visa_api/api/actions/interfaces"
	"github.com/vinaysachan/visa_api/api/request"
	"github.com/vinaysachan/visa_api/base/packages/cache"
	"github.com/vinaysachan/visa_api/data/models"
	"github.com/vinaysachan/visa_api/data/tasks"
)

type visaActionImpl struct{}

func NewVisaAction() interfaces.VisaAction {
	return &visaActionImpl{}
}

func (v *visaActionImpl) VisaApplyFormData() (*request.VisaApplyFormDataResponse, error) {

	// Use Cache Remember usage
	formData, err := cache.Remember("visa_form_data", 2*time.Hour, func() (*request.VisaApplyFormDataResponse, error) {
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

		// return countries, nil
		return &request.VisaApplyFormDataResponse{
			Countries:        countries,
			ArrivalPorts:     arrivalPorts,
			ApplicationTypes: applicationTypes,
		}, nil
	})
	if err != nil {
		return nil, errors.New("failed to remember active country list")
	}

	return formData, nil

	// cache.Set("country_list", countries, 2*time.Hour)
	// var cacheCountries []request.CountryResponse
	// cache.Get("country_list", &cacheCountries)
	// cache.Delete("country_list")
	// cache.Cleanup()

	// return &request.VisaApplyFormDataResponse{
	// 	Countries:        countries,
	// 	ArrivalPorts:     arrivalPorts,
	// 	ApplicationTypes: applicationTypes,
	// }, nil
}

func (v *visaActionImpl) VisaApplySave(req request.VisaApplicationDataRequest) (*request.VisaApplicationApplyResponse, error) {

	//Validate Nationality
	appCountry, err := tasks.FindCountryByID(req.Nationality)
	if err != nil {
		return nil, errors.New("country not found")
	}
	//Validate Port of Arrival
	arrivalPort, err := tasks.FindArrivalPortByID(req.PortOfArrival)
	if err != nil {
		return nil, errors.New("port of arrival not found")
	}
	//Validate Visa Type
	visaType, err := tasks.FindVisaTypeByID(req.VisaType)
	if err != nil {
		return nil, errors.New("visa type not found")
	}
	if visaType.ApplicationType.ID == 0 {
		return nil, errors.New("application type not found")
	}

	applicationTypeID, err := req.ApplicationType.ToUint64()
	if err != nil {
		return nil, errors.New("invalid application type")
	}

	if visaType.ApplicationType.ID != applicationTypeID {
		return nil, errors.New("visa type's Application Type not found")
	}

	application := models.AppApplication{
		ApplicationTypeID: &applicationTypeID,
		Fname:             req.Fname,
		Mname:             *req.Mname,
		Lname:             req.Lname,
	}

	if err := tasks.CreateApplication(&application); err != nil {
		return nil, errors.New("could not save Applictaion")
	}

	fmt.Println("application: ", application)

	fmt.Println("appCountry: ", appCountry, arrivalPort, visaType, visaType.ApplicationType)

	return &request.VisaApplicationApplyResponse{
		Message:       "Your e-Tourist Visa (eTV) Application submitted successfully",
		ID:            application.ID,
		ApplicationId: *application.ApplicationID,
	}, nil

}
