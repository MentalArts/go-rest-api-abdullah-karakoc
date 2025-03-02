package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorHandler manages author-related operations
type AuthorHandler struct {
	Service *services.AuthorService
}

// NewAuthorHandler, yeni bir AuthorHandler oluşturur
func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: service}
}

// GetAuthors retrieves all authors
//	@Summary		Get all authors
//	@Description	Retrieves a list of all authors
//	@Tags			authors
//	@Produce		json
//	@Success		200	{array}		dto.AuthorResponseDTO
//	@Failure		500	{object}	gin.H
//	@Router			/authors [get]
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.Service.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthor retrieves an author by ID
//	@Summary		Get an author by ID
//	@Description	Retrieves an author by their unique ID
//	@Tags			authors
//	@Produce		json
//	@Param			id	path		int	true	"Author ID"
//	@Success		200	{object}	dto.AuthorResponseDTO
//	@Failure		400	{object}	gin.H
//	@Failure		404	{object}	gin.H
//	@Router			/authors/{id} [get]
func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := h.Service.GetAuthor(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// CreateAuthor creates a new author
//	@Summary		Create a new author
//	@Description	Creates a new author using the provided details
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//	@Param			author	body		dto.CreateAuthorRequestDTO	true	"Author Data"
//	@Success		201		{object}	dto.AuthorResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/authors [post]
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var req dto.CreateAuthorRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.Service.CreateAuthor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// UpdateAuthor updates an existing author
//	@Summary		Update an author
//	@Description	Updates an existing author by ID
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Author ID"
//	@Param			author	body		dto.CreateAuthorRequestDTO	true	"Updated Author Data"
//	@Success		200		{object}	dto.AuthorResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req dto.CreateAuthorRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAuthor, err := h.Service.UpdateAuthor(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAuthor)
}

// DeleteAuthor deletes an author
//	@Summary		Delete an author
//	@Description	Deletes an author by ID
//	@Tags			authors
//	@Param			id	path	int	true	"Author ID"
//	@Success		204
//	@Failure		400	{object}	gin.H
//	@Failure		500	{object}	gin.H
//	@Router			/authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
