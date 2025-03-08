package dto

type CreateBookRequestDTO struct {
	Title           string `json:"title" binding:"required,min=3,max=200"`
	AuthorID        uint   `json:"author_id" binding:"required"`
	ISBN            string `json:"isbn" binding:"required,len=13"`
	PublicationYear int    `json:"publication_year" binding:"required,gte=1450,lte=2025"`
	Description     string `json:"description" binding:"max=500"`
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
