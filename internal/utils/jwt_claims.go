package utils

import "github.com/dgrijalva/jwt-go"

// JWTClaims represents the JWT token claims
type JWTClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
