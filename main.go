package main

import (
	"book-api-golang/config"
	"book-api-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	routes.BookRoutes(r)

	r.Run(":8080")
}
