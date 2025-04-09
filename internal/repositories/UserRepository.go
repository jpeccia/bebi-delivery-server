package repositories

import (
	"fmt"

	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetById(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user *models.User) error {
	var existingUser models.User
	err := u.db.Where("phone= ?", user.Phone).First(&existingUser).Error
	if err == nil {
		return fmt.Errorf("telefone já registrado")
	}
	return u.db.Create(user).Error
}

func (u *userRepository) Delete(id uint) error {
	return u.db.Delete(&models.User{}, id).Error
}

func (u *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user).Error
	return &user, err
}

func (u *userRepository) Update(user *models.User) error {
	return u.db.Save(user).Error
}
