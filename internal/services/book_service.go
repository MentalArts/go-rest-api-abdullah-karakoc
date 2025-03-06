package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

// BookService manages book operations
type BookService struct {
	Repo  repository.BookRepository
	Cache *redis.Client // Redis client
	Ctx   context.Context
}

// NewBookService creates a new BookService
func NewBookService(repo repository.BookRepository, cache *redis.Client, ctx context.Context) *BookService {
	return &BookService{Repo: repo, Cache: cache, Ctx: ctx}
}

// GetBooks retrieves all books and maps them to DTOs
func (s *BookService) GetBooks() ([]dto.BookResponseDTO, error) {
	// Check cache first
	cacheKey := "books_list"
	cachedData, err := s.Cache.Get(s.Ctx, cacheKey).Result()
	if err == redis.Nil { // Cache miss
		// Fetch from DB
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

		// Cache the data
		cacheData, _ := json.Marshal(bookDTOs)
		s.Cache.Set(s.Ctx, cacheKey, cacheData, 24*time.Hour) // Cache for 24 hours

		return bookDTOs, nil
	} else if err != nil {
		return nil, err
	} else {
		// Cache hit, unmarshal the cached data
		var bookDTOs []dto.BookResponseDTO
		err := json.Unmarshal([]byte(cachedData), &bookDTOs)
		if err != nil {
			return nil, err
		}
		return bookDTOs, nil
	}
}

// GetBook retrieves a specific book and maps it to a DTO
func (s *BookService) GetBook(id uint) (dto.BookResponseDTO, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("book:%d", id)
	cachedData, err := s.Cache.Get(s.Ctx, cacheKey).Result()
	if err == redis.Nil { // Cache miss
		// Fetch from DB
		book, err := s.Repo.GetBookByID(id)
		if err != nil {
			return dto.BookResponseDTO{}, err
		}

		bookDTO := dto.BookResponseDTO{
			ID:              book.ID,
			Title:           book.Title,
			ISBN:            book.ISBN,
			PublicationYear: book.PublicationYear,
			Description:     book.Description,
			AuthorID:        book.AuthorID,
			AuthorName:      book.Author.Name,
		}

		// Cache the data
		cacheData, _ := json.Marshal(bookDTO)
		s.Cache.Set(s.Ctx, cacheKey, cacheData, 24*time.Hour) // Cache for 24 hours

		return bookDTO, nil
	} else if err != nil {
		return dto.BookResponseDTO{}, err
	} else {
		// Cache hit, unmarshal the cached data
		var bookDTO dto.BookResponseDTO
		err := json.Unmarshal([]byte(cachedData), &bookDTO)
		if err != nil {
			return dto.BookResponseDTO{}, err
		}
		return bookDTO, nil
	}
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

	// Invalidate cache when creating a new book
	s.Cache.Del(s.Ctx, "books_list")

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

	// Invalidate cache when updating a book
	s.Cache.Del(s.Ctx, fmt.Sprintf("book:%d", id))
	s.Cache.Del(s.Ctx, "books_list")

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
	err := s.Repo.DeleteBook(id)
	if err != nil {
		return err
	}

	// Invalidate cache when deleting a book
	s.Cache.Del(s.Ctx, fmt.Sprintf("book:%d", id))
	s.Cache.Del(s.Ctx, "books_list")

	return nil
}
