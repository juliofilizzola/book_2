package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"net/http"
)

func LikePublication(context *gin.Context) {
	id := context.Param("id")

	var publication models.Publication

	result := initializers.DB.Find(&publication, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}
	up := initializers.DB.Model(&publication).Updates(models.Publication{
		Like: publication.Like + 1,
	})

	if up.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publication": up,
	})
}

func DislikePublication(context *gin.Context) {
	id := context.Param("id")

	var publication models.Publication

	result := initializers.DB.Find(&publication, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	if publication.Like == 0 {
		context.Status(http.StatusBadRequest)
		return
	}

	up := initializers.DB.Model(&publication).Updates(models.Publication{
		Like: publication.Like - 1,
	})

	if up.Error != nil {
		fmt.Println(result.Error)
		context.Status(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"publication": up,
	})
}
