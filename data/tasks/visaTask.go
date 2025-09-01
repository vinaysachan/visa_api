package tasks

import (
	"github.com/vinaysachan/visa_api/base/config"
	"github.com/vinaysachan/visa_api/data/models"
)

func GetActiveVisaApplicationTypes() (*[]models.AppApplicationType, error) {
	var applictaionTypes []models.AppApplicationType
	err := config.MainDB.Preload("VisaTypes", "status = ?", 1).Where("status = ?", 1).Find(&applictaionTypes).Error
	if err != nil {
		return nil, err
	}
	return &applictaionTypes, nil
}

func GetActiveVisaTypes() (*[]models.AppVisaType, error) {
	var visaTypes []models.AppVisaType
	err := config.MainDB.Where("status = ?", 1).Find(&visaTypes).Error
	if err != nil {
		return nil, err
	}
	return &visaTypes, nil
}
