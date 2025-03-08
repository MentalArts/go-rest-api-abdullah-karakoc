package dto

type CreateReviewRequestDTO struct {
	Rating     int    `json:"rating" binding:"required,gte=1,lte=5"`
	Comment    string `json:"comment" binding:"max=500"`
	DatePosted string `json:"date_posted" binding:"required,datetime=2006-01-02"`
}

type ReviewResponseDTO struct {
	ID         uint   `json:"id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
	BookID     uint   `json:"book_id"`
	BookTitle  string `json:"book_title"`
}
