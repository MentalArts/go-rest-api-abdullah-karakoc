package routes

import (
	"mentalartsapi/internal/handlers"
	"mentalartsapi/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, bookHandler *handlers.BookHandler, authorHandler *handlers.AuthorHandler, reviewHandler *handlers.ReviewHandler, authHandler *handlers.AuthHandler) {
	v1 := router.Group("/api/v1")
	{
		// Public routes for authentication
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", authHandler.RegisterUser)
			authRoutes.POST("/login", authHandler.LoginUser)
			authRoutes.POST("/refresh-token", authHandler.RefreshToken)
		}

		// Protected routes with JWT authentication
		v1.Use(middlewares.JWTAuthMiddleware()) // Protect these routes with JWT Auth Middleware

		// Book routes (Admin or Author can perform POST, PUT, DELETE)
		books := v1.Group("/books")
		{
			books.GET("/", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.POST("/", middlewares.AdminOnly(), bookHandler.CreateBook)      // Only Admin can POST
			books.PUT("/:id", middlewares.AdminOnly(), bookHandler.UpdateBook)    // Only Admin can PUT
			books.DELETE("/:id", middlewares.AdminOnly(), bookHandler.DeleteBook) // Only Admin can DELETE
			books.GET("/:id/reviews", reviewHandler.GetReviewsForBook)
			books.POST("/:id/reviews", reviewHandler.CreateReview)
		}

		// Author routes (Admin or Author can perform POST, PUT, DELETE)
		authors := v1.Group("/authors")
		{
			authors.GET("/", authorHandler.GetAuthors)
			authors.GET("/:id", authorHandler.GetAuthor)
			authors.POST("/", middlewares.AdminOnly(), authorHandler.CreateAuthor)      // Only Admin can POST
			authors.PUT("/:id", middlewares.AdminOnly(), authorHandler.UpdateAuthor)    // Only Admin can PUT
			authors.DELETE("/:id", middlewares.AdminOnly(), authorHandler.DeleteAuthor) // Only Admin can DELETE
		}

		// Review routes
		reviews := v1.Group("/reviews")
		{
			reviews.PUT("/:id", middlewares.AdminOnly(), reviewHandler.UpdateReview)    // Only Admin can PUT
			reviews.DELETE("/:id", middlewares.AdminOnly(), reviewHandler.DeleteReview) // Only Admin can DELETE
		}

	}
}
