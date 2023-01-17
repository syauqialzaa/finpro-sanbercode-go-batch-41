package config

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDatabase() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error load .env file.")
	} else {
		fmt.Println("successfully to load .env file.")
	}

	psqlConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		fmt.Println("failed to connect database.")
		panic(err)
	} else {
		fmt.Println("successfuly to connect database.")
	}

	err = db.Ping()
	database.Migrations(db)
	defer db.Close()

	return db
}