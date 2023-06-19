package auth

import (
	"github.com/gin-gonic/gin"
)

func ValidUser(context *gin.Context, userId string) bool {
	userIdToken, err := GetUserId(context)
	if err != nil {
		return false
	}

	if userId == userIdToken {
		return true
	}

	return false
}
