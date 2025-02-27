package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

// BookRepository interface for book repository
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type bookRepo struct{}

// NewBookRepository creates a new book repository
func NewBookRepository() BookRepository {
	return &bookRepo{}
}

func (r *bookRepo) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := config.DB.Preload("Author").Preload("Reviews").Find(&books).Error
	return books, err
}

func (r *bookRepo) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	err := config.DB.Preload("Author").Preload("Reviews").First(&book, id).Error
	return book, err
}

func (r *bookRepo) CreateBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

func (r *bookRepo) UpdateBook(book *models.Book) error {
	return config.DB.Save(book).Error
}

func (r *bookRepo) DeleteBook(id uint) error {
	return config.DB.Delete(&models.Book{}, id).Error
}
