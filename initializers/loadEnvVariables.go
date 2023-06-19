package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	UrlDatabase = ""
	PORT        = ""
	SecretKey   = ""
)

func LoadEnvVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT = fmt.Sprint(os.Getenv("API_PORT"))

	fmt.Println("init")
	UrlDatabase = fmt.Sprint(os.Getenv("URL_DATABASE_ENV"))

	SecretKey = fmt.Sprint(os.Getenv("SECRET_KEY"))
}
