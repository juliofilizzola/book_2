package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func UpdatePassword(context *gin.Context) {
	userIdToken, err := auth.GetUserId(context)
	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	var body struct {
		NewPassword     string `json:"new_password"`
		CurrentPassword string `json:"current_password"`
	}

	err = context.Bind(body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	id := context.Param("id")

	if userIdToken != id {
		fmt.Println("user not update")
		context.Status(http.StatusBadRequest)
		return
	}

	var password models.Password

	var user models.User

	password = models.Password{
		NewPassword:     body.NewPassword,
		CurrentPassword: body.CurrentPassword,
	}

	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	if err = auth.ValidPassword(user.Password, password.CurrentPassword); err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	newPassword, err := auth.Hash(password.NewPassword)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	up := initializers.DB.Model(&user).Where("id = ?", id).Update("password", newPassword)

	if up.Error != nil {
		fmt.Println(up.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
