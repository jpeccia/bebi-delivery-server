package services

import (
	"errors"

	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"github.com/jpeccia/bebi-delivery-server/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	user.Role = "CUSTOMER"
	return s.repo.Create(user)
}

func (s *UserService) UpgradeToStoreOwner(userID uint) error {
	user, err := s.repo.GetById(userID)
	if err != nil {
		return err
	}
	if user.Role == "STORE_OWNER" {
		return errors.New("usuário já é um StoreOwner")
	}

	user.Role = "STORE_OWNER"
	return s.repo.Update(user)
}