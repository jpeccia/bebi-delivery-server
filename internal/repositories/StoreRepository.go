package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *models.Store) error
	GetById(id uint) (*models.Store, error)
	GetAll() ([]models.Store, error)
	Update(store *models.Store) error
	Delete(id uint) error
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db}
}

func (s *storeRepository) Create(store *models.Store) error {
	return s.db.Create(store).Error
}

func (s *storeRepository) Delete(id uint) error {
	return s.db.Delete(&models.Store{}, id).Error
}

func (s *storeRepository) GetAll() ([]models.Store, error) {
	var stores []models.Store
	err := s.db.Find(&stores).Error
	return stores, err
}

func (s *storeRepository) GetById(id uint) (*models.Store, error) {
	var store models.Store
	err := s.db.First(&store, id).Error
	return &store, err
}

func (s *storeRepository) Update(store *models.Store) error {
	return s.db.Save(store).Error
}

