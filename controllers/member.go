package controllers

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/models"
	"final-project-sanbercode-go-batch-41/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMember(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()

	members, err := services.GetAllMember(db)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": members,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func GetMemberByTeamID(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()
	teamID, _ := strconv.Atoi(ctx.Param("id"))

	members, err := services.GetMemberByTeamID(db, teamID)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": members,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func AddMember(ctx *gin.Context) {
	var member models.Member
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&member)
	if err != nil {
		panic(err)
	}

	err = services.AddMember(db, member)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success create a member.",
		"result": member,
	})
}

func UpdateMember(ctx *gin.Context) {
	var member models.Member
	var db = config.ConnectToDatabase()
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&member)
	if err != nil {
		panic(err)
	}

	member.ID = id
	err = services.UpdateMember(db, member)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success update a member.",
		"result": member,
	})
}

func DeleteMember(ctx *gin.Context) {
	var member models.Member
	var db = config.ConnectToDatabase()
	id, err := strconv.Atoi(ctx.Param("id"))
	
	member.ID = id
	err = services.DeleteMember(db, member)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success delete a member.",
	})
}