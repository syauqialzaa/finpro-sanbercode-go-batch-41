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

func Register(ctx *gin.Context) {
	var admin models.Admin
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		panic(err)
	}

	if admin.Username == "" || admin.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "username and password required.",
		})
		panic(err)
	}

	// HASH PASSWORD
	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error": "failed to hash password.",
		})
		panic(err)
	}

	sql := `INSERT INTO admins (username, password) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(sql, admin.Username, password).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error": "failed to add admin to database.",
		})
		panic(err)
	}

	// GENERATE JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"username": admin.Username,
		"admin_id": admin.ID,
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
		"message": "successfully register.",
		"token": tokenString,
	})
}