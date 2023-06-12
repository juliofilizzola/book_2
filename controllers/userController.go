package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func CreateUser(context *gin.Context) {
	// @todo create user

	var body struct {
		Name     string
		Email    string
		Nick     string
		Password string
	}
	// @todo create validation email
	// @todo create password encryption
	err := context.Bind(&body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Nick:     body.Nick,
		Password: body.Password,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func GetUser(context *gin.Context) {
	var user models.User

	id := context.Param("id")
	// @todo validation from not found

	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}

func GetUsers(context *gin.Context) {
	var users []models.User
	// @todo validation from not found
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"users": users,
	})
}
