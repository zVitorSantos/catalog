package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/catalog.git/schemas"
)

// @BasePath /api/v1

// @Summary Delete product
// @Description Delete a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} DeleteProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [delete]
func DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}
	product := schemas.Product{}

	// Find Product
	if err := db.First(&product, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}

	// Delete Product
	if err := db.Delete(&product).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-product", product)

}
