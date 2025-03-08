package services

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

// RegisterUser registers a new user
func (s *AuthService) RegisterUser(dto dto.RegisterRequestDTO) (models.User, error) {
	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// Create user in the database
	user := models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: string(hashPassword),
	}

	err = s.Repo.CreateUser(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// LoginUser checks if user credentials are valid
func (s *AuthService) LoginUser(dto dto.LoginRequestDTO) (models.User, error) {
	user, err := s.Repo.GetUserByEmail(dto.Email)
	if err != nil {
		return models.User{}, err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUserByID fetches a user by ID
func (s *AuthService) GetUserByID(userID uint) (models.User, error) {
	return s.Repo.GetUserByID(userID)
}
