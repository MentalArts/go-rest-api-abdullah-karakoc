package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := config.DB.Preload("Author").Preload("Reviews").Find(&books).Error
	return books, err
}

func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	err := config.DB.Preload("Author").Preload("Reviews").First(&book, id).Error
	return book, err
}

func CreateBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

func UpdateBook(book *models.Book) error {
	return config.DB.Save(book).Error
}

func DeleteBook(id uint) error {
	return config.DB.Delete(&models.Book{}, id).Error
}
