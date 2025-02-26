package services

import (
	"mentalartsapi/internal/models"
	"mentalartsapi/internal/repository"
)

func GetAuthors() ([]models.Author, error) {
	return repository.GetAllAuthors()
}

func GetAuthor(id uint) (models.Author, error) {
	return repository.GetAuthorByID(id)
}

func CreateAuthor(author *models.Author) error {
	return repository.CreateAuthor(author)
}

func UpdateAuthor(author *models.Author) error {
	return repository.UpdateAuthor(author)
}

func DeleteAuthor(id uint) error {
	return repository.DeleteAuthor(id)
}
