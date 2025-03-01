package services

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

// AuthorService manages author operations
type AuthorService struct {
	Repo repository.AuthorRepository
}

// NewAuthorService creates a new AuthorService
func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{Repo: repo}
}

// GetAuthors retrieves all authors and converts them to DTO format
func (s *AuthorService) GetAuthors() ([]dto.AuthorResponseDTO, error) {
	authors, err := s.Repo.GetAllAuthors()
	if err != nil {
		return nil, err
	}

	var authorDTOs []dto.AuthorResponseDTO
	for _, author := range authors {
		authorDTOs = append(authorDTOs, dto.AuthorResponseDTO{
			ID:        author.ID,
			Name:      author.Name,
			Biography: author.Biography,
			BirthDate: author.BirthDate,
		})
	}

	return authorDTOs, nil
}

// GetAuthor retrieves a specific author and converts to DTO format
func (s *AuthorService) GetAuthor(id uint) (dto.AuthorResponseDTO, error) {
	author, err := s.Repo.GetAuthorByID(id)
	if err != nil {
		return dto.AuthorResponseDTO{}, err
	}

	return dto.AuthorResponseDTO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	}, nil
}

// CreateAuthor creates a new author from DTO request
func (s *AuthorService) CreateAuthor(req dto.CreateAuthorRequestDTO) (dto.AuthorResponseDTO, error) {
	author := models.Author{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDate: req.BirthDate,
	}

	err := s.Repo.CreateAuthor(&author)
	if err != nil {
		return dto.AuthorResponseDTO{}, err
	}

	return dto.AuthorResponseDTO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	}, nil
}

// UpdateAuthor updates an existing author
func (s *AuthorService) UpdateAuthor(id uint, req dto.CreateAuthorRequestDTO) (dto.AuthorResponseDTO, error) {
	author, err := s.Repo.GetAuthorByID(id)
	if err != nil {
		return dto.AuthorResponseDTO{}, err
	}

	author.Name = req.Name
	author.Biography = req.Biography
	author.BirthDate = req.BirthDate

	err = s.Repo.UpdateAuthor(&author)
	if err != nil {
		return dto.AuthorResponseDTO{}, err
	}

	return dto.AuthorResponseDTO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	}, nil
}

// DeleteAuthor deletes an author by ID
func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.Repo.DeleteAuthor(id)
}
