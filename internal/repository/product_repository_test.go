package repository

import (
	"testing"

	"github.com/davidr88/api-products/internal/domain"
)

func TestSaveProduct(t *testing.T) {
	repo := NewProductRepo()
	product := &domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		BarCode:     "123456789",
		Price:       99.99,
		Stock:       10,
	}

	err := repo.Save(product)
	if err != nil {
		t.Errorf("Failed to save product: %v", err)
	}

	if product.ID == 0 {
		t.Error("Product ID should not be 0 after save")
	}

	if product.Stock != 10 {
		t.Errorf("Expected stock to be 10, got %d", product.Stock)
	}

	savedProduct, err := repo.GetByBarCode("123456789")
	if err != nil {
		t.Errorf("Failed to get product by barcode: %v", err)
	}

	if savedProduct.Name != product.Name {
		t.Errorf("Expected product name %s, got %s", product.Name, savedProduct.Name)
	}
}
