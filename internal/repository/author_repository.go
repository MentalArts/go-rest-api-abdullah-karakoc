package repository

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/models"
)

func GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := config.DB.Preload("Books").Find(&authors).Error
	return authors, err
}

func GetAuthorByID(id uint) (models.Author, error) {
	var author models.Author
	err := config.DB.Preload("Books").First(&author, id).Error
	return author, err
}

func CreateAuthor(author *models.Author) error {
	return config.DB.Create(author).Error
}

func UpdateAuthor(author *models.Author) error {
	return config.DB.Save(author).Error
}

func DeleteAuthor(id uint) error {
	return config.DB.Delete(&models.Author{}, id).Error
}
