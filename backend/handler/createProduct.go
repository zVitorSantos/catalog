package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/catalog.git/schemas"
)

// @BasePath /api/v1

// @Summary Create product
// @Description Create a new job product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} CreateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
func CreateProductHandler(ctx *gin.Context) {
	request := CreateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product := schemas.Product{
		Ref:          request.Ref,
		Imagem:       request.Imagem,
		Descricao:    request.Descricao,
		Categoria_1:  request.Categoria_1,
		Categoria_2:  request.Categoria_2,
		Categoria_3:  request.Categoria_3,
		Complementos: request.Complementos,
		Material:     request.Material,
		Peso:         request.Peso,
		Altura:       request.Altura,
		Largura:      request.Largura,
		Comprimento:  request.Comprimento,
		Valor:        request.Valor,
		Matriz:       *request.Matriz,
		Piloto:       *request.Piloto,
		Desenho:      *request.Desenho,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.Errorf("error creating product: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating product on database")
		return
	}

	sendSuccess(ctx, "create-product", product)
}
