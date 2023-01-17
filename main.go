package main

import (
	"final-project-sanbercode-go-batch-41/config"
	"final-project-sanbercode-go-batch-41/routers"
	"fmt"
	"os"
)

func main() {
	var db = config.ConnectToDatabase()

	var router = routers.Routers(db)
	fmt.Println("server is running...")
	router.Run(":" + os.Getenv("PORT"))
}