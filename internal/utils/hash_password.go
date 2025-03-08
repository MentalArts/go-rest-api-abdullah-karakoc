package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain password using bcrypt
func HashPassword(password string) (string, error) {
	// Hashing the password with a cost factor of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares the plain password with the hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
