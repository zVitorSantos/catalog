package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/catalog.git/schemas"
)

// @BasePath /api/v1

// @Summary List all products
// @Description List all job products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} ListProductsResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func ListProductsHandler(ctx *gin.Context) {
	products := []schemas.Product{}

	if err := db.Find(&products).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing products")
		return
	}

	sendSuccess(ctx, "list-products", products)
}
