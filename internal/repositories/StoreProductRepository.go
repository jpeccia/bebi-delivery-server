package repositories

import "github.com/jpeccia/bebi-delivery-server/internal/models"

type StoreProductRepository interface {
	Create(storeProduct *models.StoreProduct) error
	GetById(id uint) (*models.StoreProduct, error)
	GetAll() ([]models.StoreProduct, error)
	Update(storeProduct *models.StoreProduct) error
	Delete(id uint) error
}