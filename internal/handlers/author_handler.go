package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorHandler, yazar işlemlerini yöneten handler yapısı
type AuthorHandler struct {
	Service *services.AuthorService
}

// NewAuthorHandler, yeni bir AuthorHandler oluşturur
func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
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

// CreateAuthor, yeni bir yazar oluşturur (DTO kullanılarak)
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

// UpdateAuthor, var olan bir yazarı günceller (DTO kullanılarak)
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

// DeleteAuthor, ID'ye göre yazarı siler
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
