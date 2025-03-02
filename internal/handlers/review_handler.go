package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReviewHandler manages review-related operations
type ReviewHandler struct {
	Service *services.ReviewService
}

// NewReviewHandler creates a new ReviewHandler instance
func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: service}
}

// GetReviewsForBook retrieves all reviews for a specific book
//	@Summary		Get reviews for a book
//	@Description	Retrieves a list of reviews for a specific book by its ID
//	@Tags			reviews
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{array}		dto.ReviewResponseDTO
//	@Failure		400	{object}	gin.H
//	@Failure		500	{object}	gin.H
//	@Router			/books/{id}/reviews [get]
func (h *ReviewHandler) GetReviewsForBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))

	reviews, err := h.Service.GetReviews(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// CreateReview creates a new review for a book
//	@Summary		Create a new review
//	@Description	Creates a new review for a specific book
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Book ID"
//	@Param			review	body		dto.CreateReviewRequestDTO	true	"Review Data"
//	@Success		201		{object}	dto.ReviewResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/books/{id}/reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var reviewDTO dto.CreateReviewRequestDTO
	if err := c.ShouldBindJSON(&reviewDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReview, err := h.Service.CreateReview(uint(bookID), reviewDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdReview)
}

// UpdateReview updates an existing review
//	@Summary		Update a review
//	@Description	Updates an existing review by its ID
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Review ID"
//	@Param			review	body		dto.CreateReviewRequestDTO	true	"Updated Review Data"
//	@Success		200		{object}	dto.ReviewResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		404		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/reviews/{id} [put]
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

// DeleteReview deletes a review
//	@Summary		Delete a review
//	@Description	Deletes a review by its ID
//	@Tags			reviews
//	@Param			id	path	int	true	"Review ID"
//	@Success		204
//	@Failure		400	{object}	gin.H
//	@Failure		404	{object}	gin.H
//	@Failure		500	{object}	gin.H
//	@Router			/reviews/{id} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID, _ := strconv.Atoi(c.Param("id"))
	err := h.Service.DeleteReview(uint(reviewID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
