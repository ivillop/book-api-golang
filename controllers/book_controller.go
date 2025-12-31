package controllers

import (
	"book-api-golang/config"
	"book-api-golang/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Judul: "Belajar Golang", Penulis: "Akbar", Tahun: 2024},
}

// GetBooks godoc
// @Summary Get semua buku
// @Tags Books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	config.DB.Model(&models.Book{}).Count(&total)

	result := config.DB.
		Limit(limit).
		Offset(offset).
		Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data buku",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total_data": total,
			"total_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	})

	c.JSON(http.StatusOK, books)
}

// GetBookById godoc
// @Summary Get buku by ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary Tambah buku
// @Tags Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary Update buku
// @Tags Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
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

// DeleteBook godoc
// @Summary Hapus buku
// @Tags Books
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{id} [delete]
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
