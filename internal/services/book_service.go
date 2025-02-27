package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// BookService manages book operations
type BookService struct {
	Repo repository.BookRepository
}

// NewBookService creates a new BookService
func NewBookService(repo repository.BookRepository) BookService {
	return BookService{Repo: repo}
}

// GetBooks retrieves all books
func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.Repo.GetAllBooks()
}

// GetBook retrieves a specific book
func (s *BookService) GetBook(id uint) (models.Book, error) {
	return s.Repo.GetBookByID(id)
}

// CreateBook creates a new book
func (s *BookService) CreateBook(book *models.Book) error {
	return s.Repo.CreateBook(book)
}

// UpdateBook updates an existing book
func (s *BookService) UpdateBook(book *models.Book) error {
	return s.Repo.UpdateBook(book)
}

// DeleteBook deletes a book by ID
func (s *BookService) DeleteBook(id uint) error {
	return s.Repo.DeleteBook(id)
}
