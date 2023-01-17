package controllers

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/models"
	"final-project-sanbercode-go-batch-41/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTask(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()

	tasks, err := services.GetAllTask(db)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": tasks,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func GetTaskByProjectID(ctx *gin.Context) {
	var result gin.H
	var db = config.ConnectToDatabase()
	projectID, _ := strconv.Atoi(ctx.Param("id"))

	tasks, err := services.GetTaskByProjectID(db, projectID)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": tasks,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateTask(ctx *gin.Context) {
	var task models.Task
	var db = config.ConnectToDatabase()

	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		panic(err)
	}

	err = services.CreateTask(db, task)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success create a task.",
		"result": task,
	})
}

func UpdateTask(ctx *gin.Context) {
	var task models.Task
	var db = config.ConnectToDatabase()
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		panic(err)
	}

	task.ID = id
	err = services.UpdateTask(db, task)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success update a task.",
		"result": task,
	})
}

func DeleteTask(ctx *gin.Context) {
	var task models.Task
	var db = config.ConnectToDatabase()
	id, err := strconv.Atoi(ctx.Param("id"))
	
	task.ID = id
	err = services.DeleteTask(db, task)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "success delete a task.",
	})
}