package routes

import (
	"github.com/gin-gonic/gin"

	"wishlist-backend/internal/handlers"
)

func RegisterRoutes(
	router *gin.Engine,
	bondHandler *handlers.BondHandler,
	wishlistHandler *handlers.WishlistHandler,
) {

	v1 := router.Group("/api/v1")

	{
		// Bonds
		v1.GET(
			"/all-bonds",
			bondHandler.GetAllBonds,
		)

		// Wishlist
		v1.POST(
			"/wishlist",
			wishlistHandler.CreateWishlist,
		)

		v1.GET(
			"/wishlist",
			wishlistHandler.GetAllWishlists,
		)

		v1.GET(
			"/wishlist/:id",
			wishlistHandler.GetWishlistByID,
		)

		v1.PUT(
			"/wishlist/:id",
			wishlistHandler.UpdateWishlist,
		)

		v1.DELETE(
			"/wishlist/:id",
			wishlistHandler.DeleteWishlist,
		)
	}
}
