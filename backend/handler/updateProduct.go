package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/catalog.git/schemas"
)

// @BasePath /api/v1

// @Summary Update product
// @Description Update a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product Identification"
// @Param product body UpdateProductRequest true "Product data to Update"
// @Success 200 {object} UpdateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [put]
func UpdateProductHandler(ctx *gin.Context) {
	request := UpdateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	// Update product
	if request.Ref != "" {
		product.Ref = request.Ref
	}
	if request.Imagem != "" {
		product.Imagem = request.Imagem
	}
	if request.Descricao != "" {
		product.Descricao = request.Descricao
	}
	if request.Categoria_1 != "" {
		product.Categoria_1 = request.Categoria_1
	}
	if request.Categoria_2 != "" {
		product.Categoria_2 = request.Categoria_2
	}
	if request.Categoria_3 != "" {
		product.Categoria_3 = request.Categoria_3
	}
	if request.Complementos != "" {
		product.Complementos = request.Complementos
	}
	if request.Material != "" {
		product.Material = request.Material
	}
	if request.Peso > 0 {
		product.Peso = request.Peso
	}
	if request.Altura > 0 {
		product.Altura = request.Altura
	}
	if request.Largura > 0 {
		product.Largura = request.Largura
	}
	if request.Comprimento > 0 {
		product.Comprimento = request.Comprimento
	}
	if request.Valor > 0 {
		product.Valor = request.Valor
	}
	if request.Matriz != nil {
		product.Matriz = *request.Matriz
	}
	if request.Piloto != nil {
		product.Piloto = *request.Piloto
	}
	if request.Desenho != nil {
		product.Desenho = *request.Desenho
	}

	// Save product

	if err := db.Save(&product).Error; err != nil {
		logger.Errorf("error updating product: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating product")
		return
	}

	sendSuccess(ctx, "update-product", product)

}
