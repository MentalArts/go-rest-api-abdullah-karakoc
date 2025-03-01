package dto

type CreateReviewRequestDTO struct {
	Rating     int    `json:"rating" binding:"required"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
	BookID     uint   `json:"book_id" binding:"required"`
}

type ReviewResponseDTO struct {
	ID         uint   `json:"id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
	BookID     uint   `json:"book_id"`
	BookTitle  string `json:"book_title"`
}
