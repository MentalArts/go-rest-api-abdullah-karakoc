package main

import (
	"mentalartsapi/config"
	"mentalartsapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
