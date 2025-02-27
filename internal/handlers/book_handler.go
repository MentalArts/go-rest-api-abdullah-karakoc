package handlers

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookHandler, kitap işlemlerini yöneten handler yapısı
type BookHandler struct {
	Service services.BookService
}

// NewBookHandler, yeni bir BookHandler oluşturur
func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// GetBooks, tüm kitapları getirir
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBook, ID'ye göre bir kitabı getirir
func (h *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.Service.GetBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook, yeni bir kitap oluşturur
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook, var olan bir kitabı günceller
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)
	if err := h.Service.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook, ID'ye göre kitabı siler
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
