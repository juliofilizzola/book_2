package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateFollowers(context *gin.Context) {
	id := context.Param("id")
	followerId := context.Param("followerId")
	formatId, err := strconv.ParseUint(followerId, 10, 64)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	var user models.User
	var follower models.User
	initializers.DB.Preload("Followers").First(&user, "id = ?", id)

	err = initializers.DB.Model(&user).Association("Followers").Append(&models.User{
		Model: gorm.Model{
			ID: uint(formatId),
		},
	})

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	err = initializers.DB.Model(&user).Where("id = ?", id).Association("Followers").Find(&follower)
	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}

func DesFollow(context *gin.Context) {
	id := context.Param("id")
	valid := auth.ValidUser(context, id)
	if !valid {
		context.Status(http.StatusUnauthorized)
		return
	}
	followerId := context.Param("followerId")
	formatId, err := strconv.ParseUint(followerId, 10, 64)

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	var user models.User
	var follower models.User
	initializers.DB.Preload("Followers").First(&user, "id = ?", id)

	err = initializers.DB.Model(&user).Association("Followers").Delete(&models.User{
		Model: gorm.Model{
			ID: uint(formatId),
		},
	})

	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}

	err = initializers.DB.Model(&user).Where("id = ?", id).Association("Followers").Find(&follower)
	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}
