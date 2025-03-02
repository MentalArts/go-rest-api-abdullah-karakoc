package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookHandler manages book-related operations
type BookHandler struct {
	Service *services.BookService
}

// NewBookHandler creates a new BookHandler instance
func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// GetBooks retrieves all books
//	@Summary		Get all books
//	@Description	Retrieves a list of all books
//	@Tags			books
//	@Produce		json
//	@Success		200	{array}		dto.BookResponseDTO
//	@Failure		500	{object}	gin.H
//	@Router			/books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBook retrieves a book by ID
//	@Summary		Get a book by ID
//	@Description	Retrieves a book by its unique ID
//	@Tags			books
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	dto.BookResponseDTO
//	@Failure		400	{object}	gin.H
//	@Failure		404	{object}	gin.H
//	@Router			/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.Service.GetBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book
//	@Summary		Create a new book
//	@Description	Creates a new book using the provided details
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			book	body		dto.CreateBookRequestDTO	true	"Book Data"
//	@Success		201		{object}	dto.BookResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Service.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook updates an existing book
//	@Summary		Update a book
//	@Description	Updates an existing book by ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Book ID"
//	@Param			book	body		dto.CreateBookRequestDTO	true	"Updated Book Data"
//	@Success		200		{object}	dto.BookResponseDTO
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var req dto.CreateBookRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Service.UpdateBook(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book
//	@Summary		Delete a book
//	@Description	Deletes a book by ID
//	@Tags			books
//	@Param			id	path	int	true	"Book ID"
//	@Success		204
//	@Failure		400	{object}	gin.H
//	@Failure		500	{object}	gin.H
//	@Router			/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.Service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
