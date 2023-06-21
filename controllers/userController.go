package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func CreateUser(context *gin.Context) {

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

	if err := user.PrepareData(false); err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
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

func UpdateUser(context *gin.Context) {
	var body struct {
		Name  string
		Email string
		Nick  string
	}
	// @todo create validation email

	err := context.Bind(&body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	id := context.Param("id")

	// @todo validation from not found

	valid := auth.ValidUser(context, id)

	if !valid {
		context.Status(http.StatusUnauthorized)
		return
	}

	var user models.User

	user = models.User{
		Name:  body.Name,
		Email: body.Email,
		Nick:  body.Nick,
	}

	if err := user.PrepareData(true); err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	initializers.DB.First(&user, id)

	result := initializers.DB.Model(&user).Updates(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	if valid := auth.ValidUser(context, id); !valid {
		context.Status(http.StatusUnauthorized)
		return
	}

	var user models.User

	result := initializers.DB.Delete(&user, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}
