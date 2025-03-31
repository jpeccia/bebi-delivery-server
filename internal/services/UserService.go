package services

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"github.com/jpeccia/bebi-delivery-server/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s * UserService) RegisterUser(user *models.User) error {
	user.Role = "CUSTOMER"
	return s.repo.Create(user)
}