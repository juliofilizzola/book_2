package main

import (
	"github.com/juliofilizzola/book_2/initializers"
	"github.com/juliofilizzola/book_2/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Publication{})
	if err != nil {
		log.Fatal(err)
	}
}
