package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/catalog.git/schemas"
)

// @BasePath /api/v1

// @Summary Show product
// @Description Show a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} ShowProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [get]
func ShowProductHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}
	product := schemas.Product{}
	if err := db.First(&product, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "product not found")
		return
	}
	sendSuccess(ctx, "product-found", product)
}
