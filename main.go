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

	// user routes
	r.POST("/user", controllers.CreateUser)
	r.GET("/user", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("user/:id", controllers.UpdateUser)
	r.DELETE("user/:id", controllers.DeleteUser)

	// follow routes
	r.PUT("followers/:id/:followerId", controllers.CreateFollowers)
	r.DELETE("followers/:id/:followerId", controllers.DesFollow)

	// publication routes
	r.POST("publication/:id", controllers.CreatePublication)
	r.GET("publication/:id", controllers.GetPublicationsByUser)
	r.GET("publication", controllers.GetPublications)
	r.PUT("publication/:id", controllers.UpdatePublication)
	r.DELETE("publication/:id", controllers.DeletePublication)

	// like publications routes
	r.GET("like/:id", controllers.LikePublication)
	r.GET("dislike/:id", controllers.DislikePublication)

	// login route
	r.POST("login", controllers.Login)

	err := r.Run(initializers.PORT)

	if err != nil {
		log.Fatal(err)
	}
}
