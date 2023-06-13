package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func CreatePublication(context *gin.Context) {
	id := context.Param("id")
	var body struct {
		Title       string
		Description string
		Content     string
	}
	// @todo create validation email
	// @todo create password encryption
	err := context.Bind(&body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	publication := models.Publication{
		Title:       body.Title,
		Description: body.Description,
		AuthId:      id,
		Content:     body.Title,
		Like:        0,
	}

	result := initializers.DB.Create(&publication)
	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"publication": publication,
	})
}
