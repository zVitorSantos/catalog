package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zVitorSantos/catalog.git/docs"
	"github.com/zVitorSantos/catalog.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.basePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/product", handler.ShowProductHandler)
		v1.POST("/product", handler.CreateProductHandler)
		v1.DELETE("/product", handler.DeleteProductHandler)
		v1.PUT("/product", handler.UpdateProductHandler)
		v1.GET("/products", handler.GetProductsHandler)
	}

	//Initializer Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
