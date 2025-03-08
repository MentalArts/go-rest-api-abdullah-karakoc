package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"mentalartsapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorHandler, yazar işlemlerini yöneten yapıdır.
type AuthorHandler struct {
	Service *services.AuthorService
}

// NewAuthorHandler, yeni bir AuthorHandler oluşturur.
func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: service}
}

// GetAuthors, tüm yazarları getirir.
//
//	@Summary		Get all authors
//	@Description	Retrieves a list of all authors
//	@Tags			authors
//	@Produce		json
//	@Success		200	{array}		dto.AuthorResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/authors [get]
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.Service.GetAuthors()
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthor, ID'ye göre bir yazarı getirir.
//
//	@Summary		Get an author by ID
//	@Description	Retrieves an author by their unique ID
//	@Tags			authors
//	@Produce		json
//	@Param			id	path		int	true	"Author ID"
//	@Success		200	{object}	dto.AuthorResponseDTO
//	@Failure		400	{object}	dto.ErrorResponseDTO
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Router			/authors/{id} [get]
func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID)
		return
	}

	author, err := h.Service.GetAuthor(uint(id))
	if err != nil {
		c.Error(utils.ErrNotFound)
		return
	}
	c.JSON(http.StatusOK, author)
}

// CreateAuthor, yeni bir yazar oluşturur.
//
//	@Summary		Create a new author
//	@Description	Creates a new author using the provided details
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//	@Param			author	body		dto.CreateAuthorRequestDTO	true	"Author Data"
//	@Success		201		{object}	dto.AuthorResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/authors [post]
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var req dto.CreateAuthorRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.ErrBadRequest)
		return
	}

	author, err := h.Service.CreateAuthor(req)
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}

	c.JSON(http.StatusCreated, author)
}

// UpdateAuthor, bir yazarı günceller.
//
//	@Summary		Update an author
//	@Description	Updates an existing author by ID
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Author ID"
//	@Param			author	body		dto.CreateAuthorRequestDTO	true	"Updated Author Data"
//	@Success		200		{object}	dto.AuthorResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID)
		return
	}

	var req dto.CreateAuthorRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.ErrBadRequest)
		return
	}

	updatedAuthor, err := h.Service.UpdateAuthor(uint(id), req)
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, updatedAuthor)
}

// DeleteAuthor, delete an author
//
//	@Summary		Delete an author
//	@Description	Deletes an author by ID
//	@Tags			authors
//	@Param			id	path	int	true	"Author ID"
//	@Success		204
//	@Failure		400	{object}	dto.ErrorResponseDTO
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID) // Invalid ID format
		return
	}

	// Check if the author exists
	_, err = h.Service.GetAuthor(uint(id))
	if err != nil {
		// If the author is not found, return a 404 Not Found error
		c.Error(utils.ErrNotFound)
		return
	}

	// Try deleting the author
	if err := h.Service.DeleteAuthor(uint(id)); err != nil {
		// If any internal server error occurs
		c.Error(utils.ErrInternal)
		return
	}

	// Successful deletion, return 204 No Content
	c.JSON(http.StatusNoContent, nil)
}
