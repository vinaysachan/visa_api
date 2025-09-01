package tasks

import (
	"github.com/vinaysachan/visa_api/base/config"
	"github.com/vinaysachan/visa_api/data/models"
)

func GetActiveCountryList() (*[]models.AppCountry, error) {
	var countries []models.AppCountry
	err := config.MainDB.Where("status = ?", 1).Find(&countries).Error
	if err != nil {
		return nil, err
	}
	return &countries, nil
}

func GetActiveArrivalPortList() (*[]models.AppArrivalPort, error) {
	var arrivalPorts []models.AppArrivalPort
	err := config.MainDB.Where("status = ?", 1).Find(&arrivalPorts).Error
	if err != nil {
		return nil, err
	}
	return &arrivalPorts, nil
}
