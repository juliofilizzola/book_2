package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/book_2/auth"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := auth.ValidToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"err": "token invalid",
			})
			context.Abort()
		}
	}
}
