package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

// AuthorRepository interface for author repository
type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
	GetAuthorByID(id uint) (models.Author, error)
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id uint) error
}

type authorRepo struct{}

// NewAuthorRepository creates a new author repository
func NewAuthorRepository() AuthorRepository {
	return &authorRepo{}
}

func (r *authorRepo) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := config.DB.Preload("Books").Find(&authors).Error
	return authors, err
}

func (r *authorRepo) GetAuthorByID(id uint) (models.Author, error) {
	var author models.Author
	err := config.DB.Preload("Books").First(&author, id).Error
	return author, err
}

func (r *authorRepo) CreateAuthor(author *models.Author) error {
	return config.DB.Create(author).Error
}

func (r *authorRepo) UpdateAuthor(author *models.Author) error {
	return config.DB.Save(author).Error
}

func (r *authorRepo) DeleteAuthor(id uint) error {
	return config.DB.Delete(&models.Author{}, id).Error
}
