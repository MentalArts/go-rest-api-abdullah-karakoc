package dto

type CreateAuthorRequestDTO struct {
	Name      string `json:"name" binding:"required,min=3,max=30"`
	Biography string `json:"biography" binding:"max=500"`
	BirthDate string `json:"birth_date" binding:"required,datetime=2006-01-02"` // YYYY-MM-DD
}

type AuthorResponseDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}
