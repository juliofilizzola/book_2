package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/initializers"
	"log"
)

//type Repository struct {
//	DB *gorm.DB
//}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
