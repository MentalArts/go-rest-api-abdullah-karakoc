package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

func GetReviews(bookID uint) ([]models.Review, error) {
	return repository.GetReviewsForBook(bookID)
}

func CreateReview(review *models.Review) error {
	return repository.CreateReview(review)
}

func UpdateReview(review *models.Review) error {
	return repository.UpdateReview(review)
}

func DeleteReview(id uint) error {
	return repository.DeleteReview(id)
}
