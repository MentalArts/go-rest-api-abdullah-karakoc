package routes

import (
	"mentalartsapi/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, bookHandler *handlers.BookHandler, authorHandler *handlers.AuthorHandler, reviewHandler *handlers.ReviewHandler) {
	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("/", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.POST("/", bookHandler.CreateBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}

		authors := v1.Group("/authors")
		{
			authors.GET("/", authorHandler.GetAuthors)
			authors.GET("/:id", authorHandler.GetAuthor)
			authors.POST("/", authorHandler.CreateAuthor)
			authors.PUT("/:id", authorHandler.UpdateAuthor)
			authors.DELETE("/:id", authorHandler.DeleteAuthor)
		}

		reviews := v1.Group("/reviews")
		{
			reviews.PUT("/:id", reviewHandler.UpdateReview)
			reviews.DELETE("/:id", reviewHandler.DeleteReview)
			reviews.GET("/books/:id/reviews", reviewHandler.GetReviewsForBook)
			reviews.POST("/books/:id/reviews", reviewHandler.CreateReview)
		}
	}
}
