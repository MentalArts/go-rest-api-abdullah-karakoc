package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReviewHandler, yorum işlemlerini yöneten handler yapısı
type ReviewHandler struct {
	Service *services.ReviewService
}

// NewReviewHandler, yeni bir ReviewHandler oluşturur
func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: service}
}

// GetReviewsForBook, belirli bir kitabın tüm yorumlarını getirir
func (h *ReviewHandler) GetReviewsForBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	reviews, err := h.Service.GetReviews(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Error() metodunu kullanıyoruz
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// CreateReview, belirli bir kitaba yeni yorum ekler
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	var reviewDTO dto.CreateReviewRequestDTO
	if err := c.ShouldBindJSON(&reviewDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviewDTO.BookID = uint(bookID) // Kitap ID'sini atama
	createdReview, err := h.Service.CreateReview(reviewDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdReview)
}

// UpdateReview, belirli bir yorumu günceller
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	reviewID, _ := strconv.Atoi(c.Param("id"))
	var reviewDTO dto.CreateReviewRequestDTO
	if err := c.ShouldBindJSON(&reviewDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedReview, err := h.Service.UpdateReview(uint(reviewID), reviewDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedReview)
}

// DeleteReview, belirli bir yorumu siler
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID, _ := strconv.Atoi(c.Param("id"))
	err := h.Service.DeleteReview(uint(reviewID)) // Hata tipi `error` olarak döndürülecek
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
