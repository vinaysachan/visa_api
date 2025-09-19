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

func FindApplicationTypeByID(id uint64) (*models.AppApplicationType, error) {
	var applicationType models.AppApplicationType
	err := config.MainDB.Where("id = ?", id).Find(&applicationType).Error
	if err != nil {
		return nil, err
	}
	return &applicationType, nil
}

func GetActiveVisaTypes() (*[]models.AppVisaType, error) {
	var visaTypes []models.AppVisaType
	err := config.MainDB.Where("status = ?", 1).Find(&visaTypes).Error
	if err != nil {
		return nil, err
	}
	return &visaTypes, nil
}

func FindVisaTypeByID(id uint64) (*models.AppVisaType, error) {
	var visaType models.AppVisaType
	err := config.MainDB.Preload("ApplicationType").Where("id = ?", id).Find(&visaType).Error
	if err != nil {
		return nil, err
	}
	return &visaType, nil
}

func CreateApplication(application *models.AppApplication) error {
	return config.MainDB.Create(application).Error
}
