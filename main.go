package main

import (
	"context"
	"mentalartsapi/config"
	"mentalartsapi/internal/handlers"
	"mentalartsapi/internal/repository"
	"mentalartsapi/internal/services"
	"mentalartsapi/middlewares"
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

	// Create a context for Redis operations (already handled in config.ConnectDatabase)
	ctx := context.Background()

	// Initialize repositories using the factory functions
	bookRepo := repository.NewBookRepository()
	authorRepo := repository.NewAuthorRepository()
	reviewRepo := repository.NewReviewRepository()

	// Initialize services with the corresponding repositories and Redis client
	bookService := services.NewBookService(bookRepo, config.Redis, ctx)
	authorService := services.NewAuthorService(authorRepo, config.Redis, ctx)
	reviewService := services.NewReviewService(reviewRepo, config.Redis, ctx)

	// Initialize handlers with the corresponding services
	bookHandler := handlers.NewBookHandler(bookService)
	authorHandler := handlers.NewAuthorHandler(authorService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	r := gin.Default()

	// Rate Limiting Middleware (1 second at most 100 request, burst = 150)
	rateLimiter := middlewares.NewRateLimiter(100, 150)
	r.Use(rateLimiter.Limit())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, bookHandler, authorHandler, reviewHandler)

	r.Run(":8000")
}
