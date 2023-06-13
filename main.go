package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/controllers"
	"github.com/juliofilizzola/book_2/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	// @todo create file from routes

	r.POST("/user", controllers.CreateUser)
	r.GET("/user", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("user/:id", controllers.UpdateUser)
	r.DELETE("user/:id", controllers.DeleteUser)
	r.PUT("followers/:id/:followerId", controllers.CreateFollowers)
	r.DELETE("followers/:id/:followerId", controllers.DesFollow)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
