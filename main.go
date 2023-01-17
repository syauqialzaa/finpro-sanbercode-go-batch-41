package main

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var db = config.ConnectToDatabase()
	var router = routers.Routers(db)

	fmt.Println("[DEBUG] successfully to load .env file.")
	fmt.Println("[DEBUG] successfuly to connect database.")
	fmt.Println("[DEBUG] sql migration file is successfully loaded.")
	fmt.Println("[DEBUG] server is running...")
	fmt.Println()

	router.Run(":" + os.Getenv("PORT"))
}