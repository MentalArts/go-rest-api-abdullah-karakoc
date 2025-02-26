package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title           string   `json:"title"`
	AuthorID        uint     `json:"author_id"`
	ISBN            string   `json:"isbn"`
	PublicationYear int      `json:"publication_year"`
	Description     string   `json:"description"`
	Author          Author   `gorm:"foreignKey:AuthorID"`
	Reviews         []Review `gorm:"foreignKey:BookID"`
}
