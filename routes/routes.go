package routes

import (
	"book-api-golang/controllers"
	"book-api-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {

	r.POST("/login", controllers.Login)

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookById)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuth())
	{
		auth.POST("/books", controllers.CreateBook)
		auth.PUT("/books/:id", controllers.UpdateBook)
		auth.DELETE("/books/:id", controllers.DeleteBook)
	}
}
