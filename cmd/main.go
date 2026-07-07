package main

import (
	"log"

	"github.com/joho/godotenv"

	"wishlist-backend/internal/config"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Database Connected Successfully 🚀")
}
