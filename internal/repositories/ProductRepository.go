package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetById(id uint) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) Create(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&models.Product{}, id).Error
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := p.db.Find(&products).Error
	return products, err
}

func (p *productRepository) GetById(id uint) (*models.Product, error) {
	var product models.Product
	err := p.db.First(&product, id).Error
	return &product, err
}

func (p *productRepository) Update(product *models.Product) error {
	return p.db.Save(product).Error
}

