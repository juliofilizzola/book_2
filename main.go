package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/controllers"
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/middlewares"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	// @todo create file from routes

	authorized := r.Group("/")

	// @todo verificar a possibilidade de criação de um group dentro de outro group
	authorized.Use(middlewares.Authentication())
	{
		// user routes
		authorized.GET("/user", controllers.GetUsers)
		authorized.GET("/user/:id", controllers.GetUser)
		authorized.PATCH("user/:id", controllers.UpdateUser)
		authorized.DELETE("user/:id", controllers.DeleteUser)

		// follow routes
		authorized.PUT("followers/:id/:followerId", controllers.CreateFollowers)
		authorized.DELETE("followers/:id/:followerId", controllers.DesFollow)

		// publication routes
		authorized.POST("publication/:id", controllers.CreatePublication)
		authorized.GET("publication/:id", controllers.GetPublicationsByUser)
		authorized.GET("publication", controllers.GetPublications)
		authorized.PUT("publication/:id", controllers.UpdatePublication)
		authorized.DELETE("publication/:id", controllers.DeletePublication)

		// like publications routes
		authorized.GET("like/:id", controllers.LikePublication)
		authorized.GET("dislike/:id", controllers.DislikePublication)

		// auth router
		authorized.PATCH("updatePassword/:id", controllers.UpdatePassword)
	}

	// user routes
	r.POST("/user", controllers.CreateUser)

	// login route
	r.POST("login", controllers.Login)

	err := r.Run(initializers.PORT)

	if err != nil {
		log.Fatal(err)
	}
}
