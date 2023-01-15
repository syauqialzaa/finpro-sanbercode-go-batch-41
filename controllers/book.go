package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	var result gin.H

	books, err := repository.GetAllBook(database.DbConnection)
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

func GetBooksByCategoryID(ctx *gin.Context) {
	var result gin.H
	category_id, _ := strconv.Atoi(ctx.Param("id"))

	books, err := repository.GetBooksByCategoryID(database.DbConnection, category_id)
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

func AddBook(ctx *gin.Context) {
	var book structs.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	match, _ := regexp.MatchString("^(http|https)://", book.ImageUrl)
	if !match {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "image_url must be inputted as a url.",
		})
		ctx.Abort()
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "release_year must be input at least 1980 and maximum 2021.",
		})
		ctx.Abort()
		return
	}

	err = repository.AddBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success inserting a book",
	})
}

func UpdateBook(ctx *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}
	
	book.ID = id

	match, _ := regexp.MatchString("^(http|https)://", book.ImageUrl)
	if !match {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "image_url harus diinput berupa url.",
		})
		ctx.Abort()
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "release_year must be input at least 1980 and maximum 2021.",
		})
		ctx.Abort()
		return
	}

	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success updating a book",
	})
}

func DeleteBook(ctx *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(ctx.Param("id"))

	book.ID = id
	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H {
		"result": "success deleting a book",
	})
}