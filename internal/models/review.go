package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
	BookID     uint   `json:"book_id"`
}
