package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := context.Bind(&body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	var user models.User

	result := initializers.DB.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	err = auth.ValidPassword(user.Password, body.Password)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
}
