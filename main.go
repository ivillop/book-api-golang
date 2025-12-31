package main

import (
	"book-api-golang/config"
	"book-api-golang/routes"

	_ "book-api-golang/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API Golang
// @version 1.0
// @description REST API Book dengan JWT Authentication
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.ConnectDatabase()
	routes.BookRoutes(r)

	r.Run(":8080")
}
