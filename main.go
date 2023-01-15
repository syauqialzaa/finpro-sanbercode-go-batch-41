package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-3/controllers"
	"quiz-3/database"
	"quiz-3/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)

func main() {
	// ENV CONFIGURATION
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load .env file.")
	} else {
		fmt.Println("Success read .env file.")
	}

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB connection is failed.")
		panic(err)
	} else {
		fmt.Println("Db connection is success.")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()
	
	// ROUTER BANGUN DATAR
	router.GET("/bangun-datar/luas-segitiga", controllers.LuasSegitigaSamaSisi)
	router.GET("/bangun-datar/keliling-segitiga", controllers.KelilingSegitigaSamaSisi)
	router.GET("/bangun-datar/luas-persegi", controllers.LuasPersegi)
	router.GET("/bangun-datar/keliling-persegi", controllers.KelilingPersegi)
	router.GET("/bangun-datar/luas-persegi-panjang", controllers.LuasPersegiPanjang)
	router.GET("/bangun-datar/keliling-persegi-panjang", controllers.KelilingPersegiPanjang)
	router.GET("/bangun-datar/luas-lingkaran", controllers.LuasLingkaran)
	router.GET("/bangun-datar/keliling-lingkaran", controllers.KelilingLingkaran)

	// ROUTER CATEGORY
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", middleware.BasicAuth, controllers.AddCategory)
	router.PUT("/categories/:id", middleware.BasicAuth, controllers.UpdateCategory)
	router.DELETE("/categories/:id", middleware.BasicAuth, controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetBooksByCategoryID)

	// ROUTER BOOK
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", middleware.BasicAuth, controllers.AddBook)
	router.PUT("books/:id", middleware.BasicAuth, controllers.UpdateBook)
	router.DELETE("books/:id", middleware.BasicAuth, controllers.DeleteBook)

	router.Run("localhost:8080")
}