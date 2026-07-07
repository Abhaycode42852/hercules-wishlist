package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"wishlist-backend/internal/config"
	"wishlist-backend/internal/handlers"
	"wishlist-backend/internal/repository"
	"wishlist-backend/internal/routes"
	"wishlist-backend/internal/services"
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

	router := gin.Default()

	// Repository
	bondRepo := repository.NewBondRepository(db)

	// Service
	bondService := services.NewBondService(bondRepo)

	// Handler
	bondHandler := handlers.NewBondHandler(bondService)

	// Wishlist Repository
	wishlistRepo :=
		repository.NewWishlistRepository(db)

	// Wishlist Service
	wishlistService :=
		services.NewWishlistService(
			wishlistRepo,
			bondRepo,
		)

	// Wishlist Handler
	wishlistHandler :=
		handlers.NewWishlistHandler(
			wishlistService,
		)

	// Routes
	routes.RegisterRoutes(
		router,
		bondHandler,
		wishlistHandler,
	)

	log.Println("Server Running on :8080")

	router.Run(":8080")
}
