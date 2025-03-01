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

// GetReviewsForBook retrieves all reviews for a given book and preloads the Book data
func (r *reviewRepo) GetReviewsForBook(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	// Preload ile ilişkili Book verisini de getiriyoruz
	err := config.DB.Preload("Book").Where("book_id = ?", bookID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetReviewByID retrieves a review by its ID, including the associated Book
func (r *reviewRepo) GetReviewByID(id uint) (models.Review, error) {
	var review models.Review
	// Preload "Book" ile ilişkili veriyi de alıyoruz
	err := config.DB.Preload("Book").First(&review, id).Error
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
