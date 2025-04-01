package handler

import (
	"net/http"

	"github.com/davidr88/api-products/internal/domain"
	"github.com/davidr88/api-products/internal/repository"
	"github.com/davidr88/api-products/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productRepo  *repository.ProductRepo
	stockService *service.MockStockService
}

func NewProductHandler(repo *repository.ProductRepo, stockService *service.MockStockService) *ProductHandler {
	return &ProductHandler{
		productRepo:  repo,
		stockService: stockService,
	}
}

func (h *ProductHandler) getProductStock(c *gin.Context, barCode string) (int, bool) {
	stockChan := make(chan int)
	errChan := make(chan error)

	go func() {
		stock, err := h.stockService.GetStock(barCode)
		if err != nil {
			errChan <- err
			return
		}
		stockChan <- stock
	}()

	select {
	case stock := <-stockChan:
		return stock, true
	case <-errChan:
		c.Status(http.StatusInternalServerError)
		return 0, false
	}
}

func (h *ProductHandler) handlePanic(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	defer h.handlePanic(c)
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err := h.productRepo.Save(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	defer h.handlePanic(c)
	products, _ := h.productRepo.GetAll()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByBarCode(c *gin.Context) {
	defer h.handlePanic(c)
	product, err := h.productRepo.GetByBarCode(c.Param("barcode"))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if stock, ok := h.getProductStock(c, product.BarCode); ok {
		product.Stock = stock
	} else {
		return
	}

	c.JSON(http.StatusOK, product)
}
