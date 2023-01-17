package main

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error load .env file.")
	} else {
		fmt.Println("successfully to load .env file.")
	}

	gin.SetMode(gin.ReleaseMode)
	var db = config.ConnectToDatabase()

	var router = routers.Routers(db)
	fmt.Println("server is running...")
	router.Run(":" + os.Getenv("PORT"))
}