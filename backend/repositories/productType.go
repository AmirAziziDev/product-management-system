package repositories

import (
	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/jmoiron/sqlx"
)

// ProductTypeRepository defines the interface for product type data operations
type ProductTypeRepository interface {
	GetProductTypes() ([]models.ProductType, error)
}

// productTypeRepository implements ProductTypeRepository
type productTypeRepository struct {
	db *sqlx.DB
}

// NewProductTypeRepository creates a new product type repository instance
func NewProductTypeRepository(db *sqlx.DB) ProductTypeRepository {
	return &productTypeRepository{db: db}
}

// GetProductTypes retrieves all product types ordered by created_at DESC
func (r *productTypeRepository) GetProductTypes() ([]models.ProductType, error) {
	var productTypes []models.ProductType
	query := "SELECT id, code, name, created_at FROM product_types ORDER BY created_at DESC"

	err := r.db.Select(&productTypes, query)
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}
