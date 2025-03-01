package dto

type CreateBookRequestDTO struct {
	Title           string `json:"title" binding:"required"`
	AuthorID        uint   `json:"author_id" binding:"required"`
	ISBN            string `json:"isbn"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
}

type BookResponseDTO struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
	AuthorID        uint   `json:"author_id"`
	AuthorName      string `json:"author_name"`
}
