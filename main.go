package main

import (
	"mentalartsapi/config"
	"mentalartsapi/internal/handlers"
	"mentalartsapi/internal/repository"
	"mentalartsapi/internal/services"
	"mentalartsapi/routes"

	_ "mentalartsapi/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Book Library Management API
//	@version		1.0
//	@description	This is a REST API for managing books, authors, and reviews.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host		localhost:8000
// @BasePath	/api/v1
// @schemes	http
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

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, bookHandler, authorHandler, reviewHandler)

	r.Run(":8000")
}
