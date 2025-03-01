package services

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// BookService manages book operations
type BookService struct {
	Repo repository.BookRepository
}

// NewBookService creates a new BookService
func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

// GetBooks retrieves all books and maps them to DTOs
func (s *BookService) GetBooks() ([]dto.BookResponseDTO, error) {
	books, err := s.Repo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var bookDTOs []dto.BookResponseDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, dto.BookResponseDTO{
			ID:              book.ID,
			Title:           book.Title,
			ISBN:            book.ISBN,
			PublicationYear: book.PublicationYear,
			Description:     book.Description,
			AuthorID:        book.AuthorID,
			AuthorName:      book.Author.Name,
		})
	}

	return bookDTOs, nil
}

// GetBook retrieves a specific book and maps it to a DTO
func (s *BookService) GetBook(id uint) (dto.BookResponseDTO, error) {
	book, err := s.Repo.GetBookByID(id)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	return dto.BookResponseDTO{
		ID:              book.ID,
		Title:           book.Title,
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
		AuthorID:        book.AuthorID,
		AuthorName:      book.Author.Name,
	}, nil
}

// CreateBook creates a new book from a DTO
func (s *BookService) CreateBook(req dto.CreateBookRequestDTO) (dto.BookResponseDTO, error) {
	book := models.Book{
		Title:           req.Title,
		AuthorID:        req.AuthorID,
		ISBN:            req.ISBN,
		PublicationYear: req.PublicationYear,
		Description:     req.Description,
	}

	err := s.Repo.CreateBook(&book)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	// Fetch the created book with author details
	createdBook, err := s.Repo.GetBookByID(book.ID)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	return dto.BookResponseDTO{
		ID:              createdBook.ID,
		Title:           createdBook.Title,
		ISBN:            createdBook.ISBN,
		PublicationYear: createdBook.PublicationYear,
		Description:     createdBook.Description,
		AuthorID:        createdBook.AuthorID,
		AuthorName:      createdBook.Author.Name,
	}, nil
}

// UpdateBook updates an existing book using a DTO
func (s *BookService) UpdateBook(id uint, req dto.CreateBookRequestDTO) (dto.BookResponseDTO, error) {
	book, err := s.Repo.GetBookByID(id)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	book.Title = req.Title
	book.AuthorID = req.AuthorID
	book.ISBN = req.ISBN
	book.PublicationYear = req.PublicationYear
	book.Description = req.Description

	err = s.Repo.UpdateBook(&book)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	// Fetch the updated book with author details
	updatedBook, err := s.Repo.GetBookByID(book.ID)
	if err != nil {
		return dto.BookResponseDTO{}, err
	}

	return dto.BookResponseDTO{
		ID:              updatedBook.ID,
		Title:           updatedBook.Title,
		ISBN:            updatedBook.ISBN,
		PublicationYear: updatedBook.PublicationYear,
		Description:     updatedBook.Description,
		AuthorID:        updatedBook.AuthorID,
		AuthorName:      updatedBook.Author.Name,
	}, nil
}

// DeleteBook deletes a book by ID
func (s *BookService) DeleteBook(id uint) error {
	return s.Repo.DeleteBook(id)
}
