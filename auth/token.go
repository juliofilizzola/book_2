package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/juliofilizzola/book_2/initializers"
	"strings"
	"time"
)

func GenerateToken(userId uint) (string, error) {

	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	fmt.Println(permission)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	s, _ := token.SignedString([]byte(initializers.SecretKey))
	return s, nil
}

func ValidToken(context *gin.Context) error {
	tokenString := getToken(context)
	fmt.Println(tokenString, "token_@")
	token, err := jwt.Parse(tokenString, getKey)

	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalid")
}

func getToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"][0]

	formatToken := strings.Split(token, " ")
	if len(formatToken) == 2 {
		return formatToken[1]
	}
	return token
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method incorrect")
	}

	return []byte(initializers.SecretKey), nil
}

func GetUserId(context *gin.Context) (string, error) {
	tokenString := getToken(context)
	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return "", err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := fmt.Sprintf("%v", permission["userId"])
		return userId, nil
	}

	return "", errors.New("token invalid")
}
