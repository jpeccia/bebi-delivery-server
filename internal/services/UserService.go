package services

import (
	"errors"

	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"github.com/jpeccia/bebi-delivery-server/internal/repositories"
	"github.com/jpeccia/bebi-delivery-server/utils"
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

func (s *UserService) AuthenticateUser(username, password string) (string, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return "", errors.New("senha incorreta")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return token, nil
}