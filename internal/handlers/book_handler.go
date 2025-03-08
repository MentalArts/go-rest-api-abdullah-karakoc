package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"mentalartsapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookHandler, kitap işlemlerini yöneten yapıdır.
type BookHandler struct {
	Service *services.BookService
}

// NewBookHandler, yeni bir BookHandler oluşturur.
func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// GetBooks, tüm kitapları getirir.
//
//	@Summary		Get all books
//	@Description	Retrieves a list of all books
//	@Tags			books
//	@Produce		json
//	@Success		200	{array}		dto.BookResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBook, ID'ye göre bir kitabı getirir.
//
//	@Summary		Get a book by ID
//	@Description	Retrieves a book by its unique ID
//	@Tags			books
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	dto.BookResponseDTO
//	@Failure		400	{object}	dto.ErrorResponseDTO
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Router			/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID)
		return
	}

	book, err := h.Service.GetBook(uint(id))
	if err != nil {
		c.Error(utils.ErrNotFound)
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook, yeni bir kitap oluşturur.
//
//	@Summary		Create a new book
//	@Description	Creates a new book using the provided details
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			book	body		dto.CreateBookRequestDTO	true	"Book Data"
//	@Success		201		{object}	dto.BookResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.ErrBadRequest)
		return
	}

	book, err := h.Service.CreateBook(req)
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook, mevcut bir kitabı günceller.
//
//	@Summary		Update a book
//	@Description	Updates an existing book by ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Book ID"
//	@Param			book	body		dto.CreateBookRequestDTO	true	"Updated Book Data"
//	@Success		200		{object}	dto.BookResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID)
		return
	}

	var req dto.CreateBookRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.ErrBadRequest)
		return
	}

	book, err := h.Service.UpdateBook(uint(id), req)
	if err != nil {
		c.Error(utils.ErrInternal)
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook, delete a book
//
//	@Summary		Delete a book
//	@Description	Deletes a book by ID
//	@Tags			books
//	@Param			id	path	int	true	"Book ID"
//	@Success		204
//	@Failure		400	{object}	dto.ErrorResponseDTO
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidID)
		return
	}

	// Veritabanında kitap olup olmadığını kontrol et
	_, err = h.Service.GetBook(uint(id))
	if err != nil {
		// Kitap bulunamadıysa, 404 Not Found döndür
		c.Error(utils.ErrNotFound)
		return
	}

	// Kitap silme işlemi
	if err := h.Service.DeleteBook(uint(id)); err != nil {
		c.Error(utils.ErrInternal)
		return
	}

	// Silme başarılı olursa 204 No Content döndür
	c.JSON(http.StatusNoContent, nil)
}
