package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetById(id uint) (*models.Order, error)
	GetAll() ([]models.Order, error)
	Update(order *models.Order) error
	Delete(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}


func (o *orderRepository) Create(order *models.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepository) Delete(id uint) error {
	return o.db.Delete(&models.Order{}, id).Error
}

func (o *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := o.db.Find(&orders).Error
	return orders, err
}

func (o *orderRepository) GetById(id uint) (*models.Order, error) {
	var order models.Order
	err := o.db.First(&order, id).Error
	return &order, err
}

func (o *orderRepository) Update(order *models.Order) error {
	return o.db.Save(order).Error
}