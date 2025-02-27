package handlers

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReviewHandler, yorum işlemlerini yöneten handler yapısı
type ReviewHandler struct {
	Service services.ReviewService
}

// NewReviewHandler, yeni bir ReviewHandler oluşturur
func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: service}
}

// GetReviewsForBook, belirli bir kitabın tüm yorumlarını getirir
func (h *ReviewHandler) GetReviewsForBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	reviews, err := h.Service.GetReviews(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// CreateReview, belirli bir kitaba yeni yorum ekler
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.BookID = uint(bookID) // Kitap ID'sini atama
	if err := h.Service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, review)
}

// UpdateReview, belirli bir yorumu günceller
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	reviewID, _ := strconv.Atoi(c.Param("id"))
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.ID = uint(reviewID) // Yorum ID'sini atama
	if err := h.Service.UpdateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// DeleteReview, belirli bir yorumu siler
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteReview(uint(reviewID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
