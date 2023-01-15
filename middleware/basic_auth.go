package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H {
					"error": "username or password cannot be empty.",
				})
				ctx.Abort()
				return
		}
		username, password, _ := ctx.Request.BasicAuth()
		if username != "admin" || password != "password" {
				if username != "editor" || password != "secret" {
						ctx.JSON(http.StatusUnauthorized, gin.H {
							"error": "username or password entered is incorrect.",
						})
						ctx.Abort()
						return
				}
		}
		ctx.Next()
}