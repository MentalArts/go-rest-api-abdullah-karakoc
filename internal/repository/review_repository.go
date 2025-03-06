package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

// ReviewRepository interface for review repository
type ReviewRepository interface {
	GetReviewsForBook(bookID uint) ([]models.Review, error)
	GetReviewByID(id uint) (models.Review, error)
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(id uint) error
}

type reviewRepo struct{}

// NewReviewRepository creates a new review repository
func NewReviewRepository() ReviewRepository {
	return &reviewRepo{}
}

func (r *reviewRepo) GetReviewsForBook(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := config.DB.Preload("Book").Where("book_id = ?", bookID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepo) GetReviewByID(id uint) (models.Review, error) {
	var review models.Review
	err := config.DB.Preload("Book").First(&review, id).Error
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (r *reviewRepo) CreateReview(review *models.Review) error {
	return config.DB.Create(review).Error
}

func (r *reviewRepo) UpdateReview(review *models.Review) error {
	return config.DB.Save(review).Error
}

func (r *reviewRepo) DeleteReview(id uint) error {
	return config.DB.Delete(&models.Review{}, id).Error
}
