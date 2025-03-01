package services

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// ReviewService, yorum işlemlerini yöneten servis yapısı
type ReviewService struct {
	Repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

// GetReviews, belirli bir kitabın yorumlarını DTO formatında döndürür
func (s *ReviewService) GetReviews(bookID uint) ([]dto.ReviewResponseDTO, error) {
	// GetReviewsForBook fonksiyonunda Preload kullanılarak Book ilişkisini de yükle
	reviews, err := s.Repo.GetReviewsForBook(bookID)
	if err != nil {
		return nil, err
	}

	var reviewDTOs []dto.ReviewResponseDTO
	for _, review := range reviews {
		// review.Book.Title'ı alabilmek için Book ilişkisini önceden yüklemelisiniz
		reviewDTOs = append(reviewDTOs, dto.ReviewResponseDTO{
			ID:         review.ID,
			Rating:     review.Rating,
			Comment:    review.Comment,
			DatePosted: review.DatePosted,
			BookID:     review.BookID,
			BookTitle:  review.Book.Title, // Artık Book ilişkisi yüklendiği için başlık alınabiliyor
		})
	}

	return reviewDTOs, nil
}

func (s *ReviewService) CreateReview(bookID uint, req dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review := models.Review{
		Rating:     req.Rating,
		Comment:    req.Comment,
		DatePosted: req.DatePosted,
		BookID:     bookID, // 📌 URL'den gelen `bookID` burada kullanılıyor!
	}

	err := s.Repo.CreateReview(&review)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	// Yeni eklenen yorumu tekrar çekiyoruz
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

// UpdateReview, var olan bir yorumu günceller ve DTO formatında döndürür
func (s *ReviewService) UpdateReview(id uint, req dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review, err := s.Repo.GetReviewByID(id)
	if err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	// Güncellenen alanları atama
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
		BookTitle:  review.Book.Title, // Artık `review.Book` ilişkilendirildi
	}, nil
}

// DeleteReview, belirli bir yorumu siler
func (s *ReviewService) DeleteReview(id uint) error {
	err := s.Repo.DeleteReview(id)
	if err != nil {
		return err
	}
	return nil
}
