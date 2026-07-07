package routes

import (
	"github.com/gin-gonic/gin"

	"wishlist-backend/internal/handlers"
)

func RegisterRoutes(
	router *gin.Engine,
	bondHandler *handlers.BondHandler,
) {

	v1 := router.Group("/api/v1")

	{
		v1.GET(
			"/all-bonds",
			bondHandler.GetAllBonds,
		)
	}
}
