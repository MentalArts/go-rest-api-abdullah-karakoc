package handlers

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorHandler, yazar işlemlerini yöneten handler yapısı
type AuthorHandler struct {
	Service services.AuthorService
}

// NewAuthorHandler, yeni bir AuthorHandler oluşturur
func NewAuthorHandler(service services.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: service}
}

// GetAuthors, tüm yazarları döndürür
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.Service.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthor, ID'ye göre bir yazarı getirir
func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := h.Service.GetAuthor(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// CreateAuthor, yeni bir yazar oluşturur
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, author)
}

// UpdateAuthor, var olan bir yazarı günceller
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.ID = uint(id)
	if err := h.Service.UpdateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}

// DeleteAuthor, ID'ye göre yazarı siler
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
