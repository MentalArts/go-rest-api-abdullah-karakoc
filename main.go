package main

import (
	"context"
	"mentalartsapi/config"
	"mentalartsapi/internal/handlers"
	"mentalartsapi/internal/middlewares"
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

//	@host		localhost:8000
//	@BasePath	/api/v1
//	@schemes	http

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	// Connect to the database and migrate models
	config.ConnectDatabase()
	config.MigrateDB()

	// Create a context for Redis operations
	ctx := context.Background()

	// Initialize repositories and services
	bookRepo := repository.NewBookRepository()
	authorRepo := repository.NewAuthorRepository()
	reviewRepo := repository.NewReviewRepository()
	userRepo := repository.NewUserRepository(config.DB)

	bookService := services.NewBookService(bookRepo, config.Redis, ctx)
	authorService := services.NewAuthorService(authorRepo, config.Redis, ctx)
	reviewService := services.NewReviewService(reviewRepo, config.Redis, ctx)
	authService := services.NewAuthService(*userRepo)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)
	authorHandler := handlers.NewAuthorHandler(authorService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	authHandler := handlers.NewAuthHandler(authService)

	// Set up the router
	r := gin.Default()

	// Rate Limiting Middleware
	rateLimiter := middlewares.NewRateLimiter(100, 150)
	r.Use(rateLimiter.Limit())
	r.Use(middlewares.ErrorHandlerMiddleware())

	// Swagger Documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up routes (using a separate routes.go file)
	routes.SetupRoutes(r, bookHandler, authorHandler, reviewHandler, authHandler)

	// Start the server
	r.Run(":8000")
}
