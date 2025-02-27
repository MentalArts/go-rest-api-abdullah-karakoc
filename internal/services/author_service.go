package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// AuthorService manages author operations
type AuthorService struct {
	Repo repository.AuthorRepository
}

// NewAuthorService creates a new AuthorService
func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return AuthorService{Repo: repo}
}

// GetAuthors retrieves all authors
func (s *AuthorService) GetAuthors() ([]models.Author, error) {
	return s.Repo.GetAllAuthors()
}

// GetAuthor retrieves a specific author
func (s *AuthorService) GetAuthor(id uint) (models.Author, error) {
	return s.Repo.GetAuthorByID(id)
}

// CreateAuthor creates a new author
func (s *AuthorService) CreateAuthor(author *models.Author) error {
	return s.Repo.CreateAuthor(author)
}

// UpdateAuthor updates an existing author
func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	return s.Repo.UpdateAuthor(author)
}

// DeleteAuthor deletes an author by ID
func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.Repo.DeleteAuthor(id)
}
