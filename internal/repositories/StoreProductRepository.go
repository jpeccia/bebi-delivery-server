package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type StoreProductRepository interface {
	Create(storeProduct *models.StoreProduct) error
	GetById(id uint) (*models.StoreProduct, error)
	GetAll() ([]models.StoreProduct, error)
	Update(storeProduct *models.StoreProduct) error
	Delete(id uint) error
}

type storeProductRepository struct {
	db *gorm.DB
}

func NewStoreProductRepository(db *gorm.DB) StoreProductRepository {
	return &storeProductRepository{db}
}

func (s *storeProductRepository) Create(storeProduct *models.StoreProduct) error {
	return s.db.Create(storeProduct).Error
}

func (s *storeProductRepository) Delete(id uint) error {
	return s.db.Delete(&models.StoreProduct{}, id).Error
}

func (s *storeProductRepository) GetAll() ([]models.StoreProduct, error) {
	var storeProducts []models.StoreProduct
	err := s.db.Find(&storeProducts).Error
	return storeProducts, err
}

func (s *storeProductRepository) GetById(id uint) (*models.StoreProduct, error) {
	var storeProduct models.StoreProduct
	err := s.db.First(&storeProduct, id).Error
	return &storeProduct, err
}

func (s *storeProductRepository) Update(storeProduct *models.StoreProduct) error {
	return s.db.Save(storeProduct).Error
}

