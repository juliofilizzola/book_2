package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	UrlDatabase = ""
	PORT        = 0
	SecretKey   = ""
)

func LoadEnvVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		PORT = 9000
	}
	fmt.Println("init")
	UrlDatabase = fmt.Sprint(os.Getenv("URL_DATABASE_ENV"))

	SecretKey = fmt.Sprint(os.Getenv("SECRET_KEY"))
}
