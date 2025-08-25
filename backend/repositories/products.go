package repositories

import (
	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/jmoiron/sqlx"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	GetProducts(page, pageSize int) ([]models.Product, error)
	GetProductsCount() (int, error)
}

// productRepository implements ProductRepository
type productRepository struct {
	db *sqlx.DB
}

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

// GetProducts retrieves paginated products ordered by created_at DESC
func (r *productRepository) GetProducts(page, pageSize int) ([]models.Product, error) {
	offset := (page - 1) * pageSize

	var products []models.Product
	query := "SELECT id, code, name, description, created_at FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2"

	err := r.db.Select(&products, query, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductsCount returns the total count of products
func (r *productRepository) GetProductsCount() (int, error) {
	var total int
	countQuery := "SELECT COUNT(*) FROM products"

	err := r.db.Get(&total, countQuery)
	if err != nil {
		return 0, err
	}

	return total, nil
}
