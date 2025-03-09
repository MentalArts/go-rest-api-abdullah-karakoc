package handlers

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/services"
	"mentalartsapi/internal/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// RegisterUser registers a new user
//
//	@Summary		Register a new user
//	@Description	This endpoint registers a new user by providing username, email, and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.RegisterRequestDTO	true	"User Registration Info"
//	@Success		201		{object}	models.User				"User Created"
//	@Failure		400		{object}	map[string]string		"Invalid input"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/auth/register [post]
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var userDTO dto.RegisterRequestDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Service.RegisterUser(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Username,
		"email": user.Email,
	})
}

// LoginUser logs in an existing user and provides a JWT
//
//	@Summary		Login a user
//	@Description	This endpoint logs in an existing user and returns a JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		dto.LoginRequestDTO	true	"User Login Info"
//	@Success		200		{object}	map[string]string	"JWT Token"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		401		{object}	map[string]string	"Invalid credentials"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/auth/login [post]
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var loginDTO dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Service.LoginUser(loginDTO)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// RefreshToken refreshes a JWT token
//
//	@Summary		Refresh JWT Token
//	@Description	This endpoint refreshes a JWT token using the current token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Success		200				{object}	map[string]string	"New JWT Token"
//	@Failure		400				{object}	map[string]string	"Invalid token format"
//	@Failure		401				{object}	map[string]string	"Invalid or expired token"
//	@Failure		500				{object}	map[string]string	"Internal server error"
//	@Router			/auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	// Remove "Bearer " prefix
	tokenString = tokenString[7:]

	claims := &utils.JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Token is valid, generate a new one
	user, err := h.Service.GetUserByID(claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Helper function to generate JWT
func generateJWT(user models.User) (string, error) {
	claims := &utils.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			Issuer:    "your-app-name",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}
