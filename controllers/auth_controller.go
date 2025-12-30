package controllers

import (
	"book-api-golang/helpers"
	"book-api-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}
var userID uint = 1

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := helpers.HashPassword(input.Password)

	user := models.User{
		ID:       userID,
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	users = append(users, user)
	userID++

	c.JSON(http.StatusCreated, gin.H{"message": "Register berhasil"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Email == input.Email &&
			helpers.CheckPassword(user.Password, input.Password) {
			token, _ := helpers.GenerateToken(user.ID, user.Email)

			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Email atau password salah",
	})
}
