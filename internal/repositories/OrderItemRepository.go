package repositories

import "github.com/jpeccia/bebi-delivery-server/internal/models"

type OrderItemRepository interface {
	Create(orderItem *models.OrderItem) error
	GetById(id uint) (*models.OrderItem, error)
	GetAll() ([]models.OrderItem, error)
	Update(orderItem *models.OrderItem) error
	Delete(id uint) error
}
