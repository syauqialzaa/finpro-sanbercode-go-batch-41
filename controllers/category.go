package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(ctx *gin.Context) {
	var result gin.H

	books, err := repository.GetAllCategory(database.DbConnection)
	if err != nil {
		result = gin.H {
			"result": err,
		}
	} else {
		result = gin.H {
			"result": books,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func AddCategory(ctx *gin.Context) {
	var category structs.Category

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.AddCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success inserting a category",
	})
}

func UpdateCategory(ctx *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success updating a category",
	})
}

func DeleteCategory(ctx *gin.Context) {
	var category structs.Category
	id, err := strconv.Atoi(ctx.Param("id"))

	category.ID = id
	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success deleting a category",
	})
}