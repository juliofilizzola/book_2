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
