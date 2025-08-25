package tasks

import (
	"github.com/vinaysachan/visa_api/base/config"
	"github.com/vinaysachan/visa_api/base/utils"
	"github.com/vinaysachan/visa_api/data/models"
)

func IsEmailOrMobileUserExist(master_id uint64, email string, mobile utils.StringOrNumber) (bool, error) {
	var count int64
	err := config.MainDB.Model(&models.User{}).
		Where("`master_id` = ?", master_id).
		Where("email = ? OR mobile = ?", email, mobile).
		Count(&count).Error
	return count > 0, err
}

func FindActiveUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.MainDB.Where("email = ? AND status = 'Y'", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return config.MainDB.Create(user).Error
}

func FindUserByID(id uint64) (*models.User, error) {
	var user models.User
	err := config.MainDB.First(&user, id).Error
	return &user, err
}
