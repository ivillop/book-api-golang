package controllers

import (
	"book-api-golang/config"
	"book-api-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Judul: "Belajar Golang", Penulis: "Akbar", Tahun: 2024},
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}
