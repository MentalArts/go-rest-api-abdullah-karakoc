package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

// ReviewRepository interface for review repository
type ReviewRepository interface {
	GetReviewsForBook(bookID uint) ([]models.Review, error)
	GetReviewByID(id uint) (models.Review, error) // 📌 Yeni metod eklendi
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(id uint) error
}

type reviewRepo struct{}

// NewReviewRepository creates a new review repository
func NewReviewRepository() ReviewRepository {
	return &reviewRepo{}
}

// GetReviewsForBook retrieves all reviews for a given book
func (r *reviewRepo) GetReviewsForBook(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := config.DB.Where("book_id = ?", bookID).Find(&reviews).Error
	return reviews, err
}

// GetReviewByID retrieves a review by its ID, including the associated Book
func (r *reviewRepo) GetReviewByID(id uint) (models.Review, error) {
	var review models.Review
	err := config.DB.Preload("Book").First(&review, id).Error // 📌 Preload "Book" ile ilişkili veriyi de alıyoruz
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

// CreateReview creates a new review in the database
func (r *reviewRepo) CreateReview(review *models.Review) error {
	return config.DB.Create(review).Error
}

// UpdateReview updates an existing review in the database
func (r *reviewRepo) UpdateReview(review *models.Review) error {
	return config.DB.Save(review).Error
}

// DeleteReview deletes a review by its ID
func (r *reviewRepo) DeleteReview(id uint) error {
	return config.DB.Delete(&models.Review{}, id).Error
}
