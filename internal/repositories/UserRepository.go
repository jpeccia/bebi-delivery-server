package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/database"
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(userID uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
