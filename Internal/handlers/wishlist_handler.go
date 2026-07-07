package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"wishlist-backend/internal/services"
)

type WishlistHandler struct {
	service *services.WishlistService
}

func NewWishlistHandler(
	service *services.WishlistService,
) *WishlistHandler {
	return &WishlistHandler{
		service: service,
	}
}

func (h *WishlistHandler) CreateWishlist(
	c *gin.Context,
) {

	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	wishlist, err := h.service.CreateWishlist(
		req.Name,
	)

	if err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		wishlist,
	)
}

func (h *WishlistHandler) GetAllWishlists(
	c *gin.Context,
) {

	wishlists, err :=
		h.service.GetAllWishlists()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		wishlists,
	)
}

func (h *WishlistHandler) GetWishlistByID(
	c *gin.Context,
) {

	id := c.Param("id")

	wishlist, err :=
		h.service.GetWishlistByID(id)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "wishlist not found",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		wishlist,
	)
}

func (h *WishlistHandler) UpdateWishlist(
	c *gin.Context,
) {

	id := c.Param("id")

	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	err := h.service.UpdateWishlist(
		id,
		req.Name,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "wishlist updated",
		},
	)
}

func (h *WishlistHandler) DeleteWishlist(
	c *gin.Context,
) {

	id := c.Param("id")

	err := h.service.DeleteWishlist(id)

	if err != nil {

		if err.Error() == "wishlist not found" {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "wishlist deleted successfully",
		},
	)
}
