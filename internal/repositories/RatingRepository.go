package repositories

import (
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(rating *models.Rating) error
	GetById(id uint) (*models.Rating, error)
	GetAll() ([]models.Rating, error)
	Update(rating *models.Rating) error
	Delete(id uint) error
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{db}
}

func (r *ratingRepository) Create(rating *models.Rating) error {
	return r.db.Create(rating).Error
}

func (r *ratingRepository) Delete(id uint) error {
	return r.db.Delete(&models.Rating{}, id).Error
}

func (r *ratingRepository) GetAll() ([]models.Rating, error) {
	var ratings []models.Rating
	err := r.db.Find(&ratings).Error
	return ratings, err
}

func (r *ratingRepository) GetById(id uint) (*models.Rating, error) {
	var rating models.Rating
	err := r.db.First(&rating, id).Error
	return &rating, err
}

func (r *ratingRepository) Update(rating *models.Rating) error {
	return r.db.Save(rating).Error
}

