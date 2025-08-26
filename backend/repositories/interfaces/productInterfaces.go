package interfaces

import (
	"context"
	"errors"

	"github.com/AmirAziziDev/product-management-system/models"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	ListProducts(page, pageSize int) ([]models.Product, error)
	GetProductsCount() (int, error)
	CreateProduct(ctx context.Context, p models.Product, colorIDs []int) (int, error)
}

var (
	ErrProductTypeNotFound = errors.New("product_type_id not found")
	ErrColorsNotFound      = errors.New("product_color_ids not found")
)
