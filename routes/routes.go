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

		// Book routes
		books := v1.Group("/books")
		{
			books.GET("/", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.POST("/", bookHandler.CreateBook) // Only Admin or Author can post
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
			books.GET("/:id/reviews", reviewHandler.GetReviewsForBook)
			books.POST("/:id/reviews", reviewHandler.CreateReview)
		}

		// Author routes
		authors := v1.Group("/authors")
		{
			authors.GET("/", authorHandler.GetAuthors)
			authors.GET("/:id", authorHandler.GetAuthor)
			authors.POST("/", authorHandler.CreateAuthor)
			authors.PUT("/:id", authorHandler.UpdateAuthor)
			authors.DELETE("/:id", authorHandler.DeleteAuthor)
		}

		// Review routes
		reviews := v1.Group("/reviews")
		{
			reviews.PUT("/:id", reviewHandler.UpdateReview)
			reviews.DELETE("/:id", reviewHandler.DeleteReview)
		}

		// Admin Only Routes (for admin only actions)
		adminRoutes := v1.Group("/admin")
		adminRoutes.Use(middlewares.AdminOnly()) // Use the AdminOnly middleware for this group
		{
			adminRoutes.DELETE("/authors/:id", authorHandler.DeleteAuthor)
		}
	}
}
