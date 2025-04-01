package repository

import (
	"errors"

	"github.com/davidr88/api-products/internal/domain"
)

var (
	products []domain.Product
	lastID   int
)

type ProductRepo struct{}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (r *ProductRepo) Save(product *domain.Product) error {
	if product.BarCode == "" {
		return errors.New("bar code cannot be empty")
	}

	// Verificar si ya existe un producto con el mismo código de barras
	for _, p := range products {
		if p.BarCode == product.BarCode {
			return errors.New("product with this bar code already exists")
		}
	}

	lastID++
	product.ID = lastID
	products = append(products, *product)
	return nil
}

func (r *ProductRepo) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *ProductRepo) GetByBarCode(barCode string) (*domain.Product, error) {
	for _, p := range products {
		if p.BarCode == barCode {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}
