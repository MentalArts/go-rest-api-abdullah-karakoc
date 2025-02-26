package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
	Books     []Book `gorm:"foreignKey:AuthorID"`
}
