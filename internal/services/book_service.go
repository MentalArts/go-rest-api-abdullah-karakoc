package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

func GetBooks() ([]models.Book, error) {
	return repository.GetAllBooks()
}

func GetBook(id uint) (models.Book, error) {
	return repository.GetBookByID(id)
}

func CreateBook(book *models.Book) error {
	return repository.CreateBook(book)
}

func UpdateBook(book *models.Book) error {
	return repository.UpdateBook(book)
}

func DeleteBook(id uint) error {
	return repository.DeleteBook(id)
}
