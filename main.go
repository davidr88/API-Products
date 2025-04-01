package main

import (
	"github.com/davidr88/api-products/internal/handler"
	"github.com/davidr88/api-products/internal/repository"
	"github.com/davidr88/api-products/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	productRepo := repository.NewProductRepo()
	stockService := service.NewMockStockService()
	productHandler := handler.NewProductHandler(productRepo, stockService)
	router := gin.Default()

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/:barcode", productHandler.GetProductByBarCode)

	router.Run(":8080")
}
