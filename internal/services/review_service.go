package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// ReviewService, yorum işlemlerini yöneten servis yapısı
type ReviewService struct {
	Repo repository.ReviewRepository
}

// NewReviewService, yeni bir ReviewService oluşturur
func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return ReviewService{Repo: repo}
}

// GetReviews, belirli bir kitabın yorumlarını döndürür
func (s *ReviewService) GetReviews(bookID uint) ([]models.Review, error) {
	return s.Repo.GetReviewsForBook(bookID)
}

// CreateReview, yeni bir yorum ekler
func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.Repo.CreateReview(review)
}

// UpdateReview, var olan bir yorumu günceller
func (s *ReviewService) UpdateReview(review *models.Review) error {
	return s.Repo.UpdateReview(review)
}

// DeleteReview, belirli bir yorumu siler
func (s *ReviewService) DeleteReview(id uint) error {
	return s.Repo.DeleteReview(id)
}
