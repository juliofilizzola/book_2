package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
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

func GetPublicationsByUser(context *gin.Context) {
	var publications []models.Publication

	id := context.Param("id")
	// @todo validation from not found

	result := initializers.DB.Where("auth_id = ?", id).Find(&publications)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publications": publications,
	})
}

func GetPublications(context *gin.Context) {
	var publications []models.Publication

	// @todo validation from not found

	result := initializers.DB.Find(&publications)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publications": publications,
	})
}

func UpdatePublication(context *gin.Context) {
	var body struct {
		Title       string
		Content     string
		Description string
	}

	err := context.Bind(&body)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	id := context.Param("id")

	publication := models.Publication{
		Title:       body.Title,
		Description: body.Description,
		Content:     body.Title,
	}

	result := initializers.DB.First(&publication, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	valid := auth.ValidUser(context, publication.AuthId)

	if !valid {
		context.Status(http.StatusUnauthorized)
		return
	}

	up := initializers.DB.Model(&publication).Updates(models.Publication{
		Title:       body.Title,
		Description: body.Description,
		Content:     body.Content,
	})

	if up.Error != nil {
		fmt.Println(up.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publication": publication,
	})
}

func DeletePublication(context *gin.Context) {
	id := context.Param("id")

	var publication models.Publication

	result := initializers.DB.First(&publication, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	valid := auth.ValidUser(context, publication.AuthId)

	if !valid {
		context.Status(http.StatusUnauthorized)
		return
	}

	if deleted := initializers.DB.Delete(&publication, id); deleted.Error != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publication": publication,
	})
}
