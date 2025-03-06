package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

// ReviewService manages book operations
type ReviewService struct {
	Repo  repository.ReviewRepository
	Cache *redis.Client // Redis client
	Ctx   context.Context
}

// NewReviewService creates a new ReviewService
func NewReviewService(repo repository.ReviewRepository, cache *redis.Client, ctx context.Context) *ReviewService {
	return &ReviewService{Repo: repo, Cache: cache, Ctx: ctx}
}

// GetReviews retrieves all reviews for a book and maps them to DTOs
func (s *ReviewService) GetReviews(bookID uint) ([]dto.ReviewResponseDTO, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("reviews_book:%d", bookID)
	cachedData, err := s.Cache.Get(s.Ctx, cacheKey).Result()
	if err == redis.Nil { // Cache miss
		// Fetch from DB
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

		// Cache the data
		cacheData, _ := json.Marshal(reviewDTOs)
		s.Cache.Set(s.Ctx, cacheKey, cacheData, 24*time.Hour) // Cache for 24 hours

		return reviewDTOs, nil
	} else if err != nil {
		return nil, err
	} else {
		// Cache hit, unmarshal the cached data
		var reviewDTOs []dto.ReviewResponseDTO
		err := json.Unmarshal([]byte(cachedData), &reviewDTOs)
		if err != nil {
			return nil, err
		}
		return reviewDTOs, nil
	}
}

// CreateReview creates a new review for a book from a DTO
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

	// Fetch the created review with book details
	review, err = s.Repo.GetReviewByID(review.ID)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	// Invalidate cache when creating a new review
	s.Cache.Del(s.Ctx, fmt.Sprintf("reviews_book:%d", bookID))

	return dto.ReviewResponseDTO{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
		BookTitle:  review.Book.Title,
	}, nil
}

// UpdateReview updates an existing review using a DTO
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

	// Invalidate cache when updating a review
	s.Cache.Del(s.Ctx, fmt.Sprintf("reviews_book:%d", review.BookID))

	return dto.ReviewResponseDTO{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
		BookTitle:  review.Book.Title,
	}, nil
}

// DeleteReview deletes a review by ID
func (s *ReviewService) DeleteReview(id uint) error {
	review, err := s.Repo.GetReviewByID(id)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteReview(id)
	if err != nil {
		return err
	}

	// Invalidate cache when deleting a review
	s.Cache.Del(s.Ctx, fmt.Sprintf("reviews_book:%d", review.BookID))

	return nil
}
