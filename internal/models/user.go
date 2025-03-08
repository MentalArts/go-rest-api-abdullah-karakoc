package models

import (
	"gorm.io/gorm"
)

// User model for user authentication and registration
type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-"`
	Email    string `json:"email" gorm:"unique;not null"`
	Role     string `json:"role"` // Admin, User

}
