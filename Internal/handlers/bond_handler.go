package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wishlist-backend/internal/services"
)

type BondHandler struct {
	service *services.BondService
}

func NewBondHandler(
	service *services.BondService,
) *BondHandler {
	return &BondHandler{
		service: service,
	}
}

func (h *BondHandler) GetAllBonds(
	c *gin.Context,
) {

	page, _ := strconv.Atoi(
		c.DefaultQuery("page", "1"),
	)

	limit, _ := strconv.Atoi(
		c.DefaultQuery("limit", "10"),
	)

	sort := c.DefaultQuery(
		"sort",
		"name",
	)

	order := c.DefaultQuery(
		"order",
		"asc",
	)

	bonds, err := h.service.GetAllBonds(
		page,
		limit,
		sort,
		order,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	total, err := h.service.GetBondCount()

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
			"page":  page,
			"limit": limit,
			"data":  bonds,
			"total": total,
		},
	)
}
