package providers

import (
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/AmirAziziDev/product-management-system/repositories/interfaces"
	"github.com/jmoiron/sqlx"
)

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return repositories.NewProductRepository(db)
}

// NewProductTypeRepository creates a new product type repository instance
func NewProductTypeRepository(db *sqlx.DB) repositories.ProductTypeRepository {
	return repositories.NewProductTypeRepository(db)
}

// NewColorRepository creates a new color repository instance
func NewColorRepository(db *sqlx.DB) repositories.ColorRepository {
	return repositories.NewColorRepository(db)
}
