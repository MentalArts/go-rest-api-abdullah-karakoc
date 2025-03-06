package services

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// ReviewService manages book operations
type ReviewService struct {
	Repo repository.ReviewRepository
}

// NewReviewService creates a new BookService
func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) GetReviews(bookID uint) ([]dto.ReviewResponseDTO, error) {
	reviews, err := s.Repo.GetReviewsForBook(bookID)
	if err != nil {
		return nil, err
	}

	var reviewDTOs []dto.ReviewResponseDTO
	for _, review := range reviews {
		reviewDTOs = append(reviewDTOs, dto.ReviewResponseDTO{
			ID:         review.ID,
			Rating:     review.Rating,
			Comment:    review.Comment,
			DatePosted: review.DatePosted,
			BookID:     review.BookID,
			BookTitle:  review.Book.Title,
		})
	}

	return reviewDTOs, nil
}

func (s *ReviewService) CreateReview(bookID uint, req dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review := models.Review{
		Rating:     req.Rating,
		Comment:    req.Comment,
		DatePosted: req.DatePosted,
		BookID:     bookID,
	}

	err := s.Repo.CreateReview(&review)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	review, err = s.Repo.GetReviewByID(review.ID)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	return dto.ReviewResponseDTO{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
		BookTitle:  review.Book.Title,
	}, nil
}

func (s *ReviewService) UpdateReview(id uint, req dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review, err := s.Repo.GetReviewByID(id)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	review.Rating = req.Rating
	review.Comment = req.Comment

	review.DatePosted = req.DatePosted

	err = s.Repo.UpdateReview(&review)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	return dto.ReviewResponseDTO{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
		BookTitle:  review.Book.Title,
	}, nil
}

func (s *ReviewService) DeleteReview(id uint) error {
	err := s.Repo.DeleteReview(id)
	if err != nil {
		return err
	}
	return nil
}
