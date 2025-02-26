package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

func GetReviewsForBook(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := config.DB.Where("book_id = ?", bookID).Find(&reviews).Error
	return reviews, err
}

func CreateReview(review *models.Review) error {
	return config.DB.Create(review).Error
}

func UpdateReview(review *models.Review) error {
	return config.DB.Save(review).Error
}

func DeleteReview(id uint) error {
	return config.DB.Delete(&models.Review{}, id).Error
}
