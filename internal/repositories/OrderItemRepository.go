package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(orderItem *models.OrderItem) error
	GetById(id uint) (*models.OrderItem, error)
	GetAll() ([]models.OrderItem, error)
	Update(orderItem *models.OrderItem) error
	Delete(id uint) error
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (o *orderItemRepository) Create(orderItem *models.OrderItem) error {
	return o.db.Create(orderItem).Error
}

func (o *orderItemRepository) Delete(id uint) error {
	return o.db.Delete(&models.OrderItem{}, id).Error
}

func (o *orderItemRepository) GetAll() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := o.db.Find(&orderItems).Error
	return orderItems, err
}

func (o *orderItemRepository) GetById(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := o.db.Find(&orderItem).Error
	return &orderItem, err
}

func (o *orderItemRepository) Update(orderItem *models.OrderItem) error {
	return o.db.Save(orderItem).Error
}

