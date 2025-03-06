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

// AuthorService manages author operations
type AuthorService struct {
	Repo  repository.AuthorRepository
	Cache *redis.Client // Redis client
	Ctx   context.Context
}

// NewAuthorService creates a new AuthorService
func NewAuthorService(repo repository.AuthorRepository, cache *redis.Client, ctx context.Context) *AuthorService {
	return &AuthorService{Repo: repo, Cache: cache, Ctx: ctx}
}

// GetAuthors retrieves all authors and converts them to DTO format
func (s *AuthorService) GetAuthors() ([]dto.AuthorResponseDTO, error) {
	// Check cache first
	cacheKey := "authors_list"
	cachedData, err := s.Cache.Get(s.Ctx, cacheKey).Result()
	if err == redis.Nil { // Cache miss
		// Fetch from DB
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

		// Cache the data
		cacheData, _ := json.Marshal(authorDTOs)
		s.Cache.Set(s.Ctx, cacheKey, cacheData, 24*time.Hour) // Cache for 24 hours

		return authorDTOs, nil
	} else if err != nil {
		return nil, err
	} else {
		// Cache hit, unmarshal the cached data
		var authorDTOs []dto.AuthorResponseDTO
		err := json.Unmarshal([]byte(cachedData), &authorDTOs)
		if err != nil {
			return nil, err
		}
		return authorDTOs, nil
	}
}

// GetAuthor retrieves a specific author and converts to DTO format
func (s *AuthorService) GetAuthor(id uint) (dto.AuthorResponseDTO, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("author:%d", id)
	cachedData, err := s.Cache.Get(s.Ctx, cacheKey).Result()
	if err == redis.Nil { // Cache miss
		// Fetch from DB
		author, err := s.Repo.GetAuthorByID(id)
		if err != nil {
			return dto.AuthorResponseDTO{}, err
		}

		authorDTO := dto.AuthorResponseDTO{
			ID:        author.ID,
			Name:      author.Name,
			Biography: author.Biography,
			BirthDate: author.BirthDate,
		}

		// Cache the data
		cacheData, _ := json.Marshal(authorDTO)
		s.Cache.Set(s.Ctx, cacheKey, cacheData, 24*time.Hour) // Cache for 24 hours

		return authorDTO, nil
	} else if err != nil {
		return dto.AuthorResponseDTO{}, err
	} else {
		// Cache hit, unmarshal the cached data
		var authorDTO dto.AuthorResponseDTO
		err := json.Unmarshal([]byte(cachedData), &authorDTO)
		if err != nil {
			return dto.AuthorResponseDTO{}, err
		}
		return authorDTO, nil
	}
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

	// Invalidate cache when creating a new author
	s.Cache.Del(s.Ctx, "authors_list")

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

	// Invalidate cache when updating an author
	s.Cache.Del(s.Ctx, fmt.Sprintf("author:%d", id))
	s.Cache.Del(s.Ctx, "authors_list")

	return dto.AuthorResponseDTO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	}, nil
}

// DeleteAuthor deletes an author by ID
func (s *AuthorService) DeleteAuthor(id uint) error {
	err := s.Repo.DeleteAuthor(id)
	if err != nil {
		return err
	}

	// Invalidate cache when deleting an author
	s.Cache.Del(s.Ctx, fmt.Sprintf("author:%d", id))
	s.Cache.Del(s.Ctx, "authors_list")

	return nil
}
