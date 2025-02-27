package main

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/handlers"
	"mentalartsapi/internal/repository"
	"mentalartsapi/internal/services"
	"mentalartsapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()
	config.MigrateDB()

	// Initialize repositories using the factory functions
	bookRepo := repository.NewBookRepository()
	authorRepo := repository.NewAuthorRepository()
	reviewRepo := repository.NewReviewRepository()

	// Initialize services with the corresponding repositories
	bookService := services.NewBookService(bookRepo)
	authorService := services.NewAuthorService(authorRepo)
	reviewService := services.NewReviewService(reviewRepo)

	// Initialize handlers with the corresponding services
	bookHandler := handlers.NewBookHandler(bookService)
	authorHandler := handlers.NewAuthorHandler(authorService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// Create a new Gin router
	r := gin.Default()

	// Setup API routes with the initialized handlers
	routes.SetupRoutes(r, bookHandler, authorHandler, reviewHandler)

	// Start the server on port 8080
	r.Run(":8000")
}
