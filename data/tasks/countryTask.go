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

func FindCountryByID(id uint64) (*models.AppCountry, error) {
	var country models.AppCountry
	err := config.MainDB.First(&country, id).Error
	return &country, err
}

func GetActiveArrivalPortList() (*[]models.AppArrivalPort, error) {
	var arrivalPorts []models.AppArrivalPort
	err := config.MainDB.Where("status = ?", 1).Find(&arrivalPorts).Error
	if err != nil {
		return nil, err
	}
	return &arrivalPorts, nil
}

func FindArrivalPortByID(id uint64) (*models.AppArrivalPort, error) {
	var arrival_port models.AppArrivalPort
	err := config.MainDB.First(&arrival_port, id).Error
	return &arrival_port, err
}
