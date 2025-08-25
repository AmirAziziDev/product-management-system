package providers

import (
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/jmoiron/sqlx"
)

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *sqlx.DB) repositories.ProductRepository {
	return repositories.NewProductRepository(db)
}

// NewProductTypeRepository creates a new product type repository instance
func NewProductTypeRepository(db *sqlx.DB) repositories.ProductTypeRepository {
	return repositories.NewProductTypeRepository(db)
}
