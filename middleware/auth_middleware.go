package middleware

import (
	"final-project-sanbercode-go-batch-41/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var admin models.Admin
		tokenString := ctx.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("SecretToken"), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H { 
				"error": "invalid token.",
			})
			ctx.Abort()
			panic(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set(admin.Username, claims[admin.Username])
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H { 
				"error": "invalid token.",
			})
			ctx.Abort()
			panic(err)
		}

		ctx.Next()
	}
}