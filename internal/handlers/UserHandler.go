package handlers

import "github.com/jpeccia/bebi-delivery-server/internal/services"

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

