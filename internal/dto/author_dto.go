package dto

type CreateAuthorRequestDTO struct {
	Name      string `json:"name" binding:"required"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

type AuthorResponseDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}
