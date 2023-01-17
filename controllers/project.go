package controllers

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/models"
	"final-project-sanbercode-go-batch-41/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetAllProject(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()

	projects, err := services.GetAllProject(db)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": projects,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateProject(ctx *gin.Context) {
	var project models.Project
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		panic(err)
	}

	err = services.CreateProject(db, project)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success create a project.",
		"result": project,
	})
}

func UpdateProject(ctx *gin.Context) {
	var project models.Project
	var db = config.ConnectToDatabase()
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		panic(err)
	}

	project.ID = id
	err = services.UpdateProject(db, project)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success update a project.",
		"result": project,
	})
}

func DeleteProject(ctx *gin.Context) {
	var project models.Project
	var db = config.ConnectToDatabase()
	id, err := strconv.Atoi(ctx.Param("id"))
	
	project.ID = id
	err = services.DeleteProject(db, project)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success delete a project.",
	})
}