package middleware

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var admin models.Admin
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		panic(err)
	}

	// CHECK DATA ADMIN
	var dbAdmin models.Admin
	err = db.QueryRow(
		`SELECT id, username, password FROM admins WHERE username = $1`, admin.Username).Scan(
			&dbAdmin.ID, &dbAdmin.Username, &dbAdmin.Password,
		)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H {
			"error": "invalid username or password.",
		})
		panic(err)
	}

	// CHECK PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password))
	if err != nil {
    ctx.JSON(http.StatusUnauthorized, gin.H {
			"error": "invalid username or password.",
		})
    panic(err)
	}

	// GENERATE JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": dbAdmin.Username,
    "admin_id": dbAdmin.ID,
    "exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// FIRM TOKEN
	tokenString, err := token.SignedString([]byte("SecretToken"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error": "failed to generate token.",
		})
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "successfully login.",
		"token": tokenString,
	})
}