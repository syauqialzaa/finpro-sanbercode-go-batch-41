package controllers

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/models"
	"final-project-sanbercode-go-batch-41/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTeam(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()

	teams, err := services.GetAllTeam(db)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": teams,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateTeam(ctx *gin.Context) {
	var team models.Team
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		panic(err)
	}

	err = services.CreateTeam(db, team)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success create a team.",
		"result": team,
	})
}