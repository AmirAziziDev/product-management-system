package repositories

import (
	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/AmirAziziDev/product-management-system/repositories/interfaces"
	"github.com/jmoiron/sqlx"
)

// productRepository implements ProductRepository
type productRepository struct {
	db *sqlx.DB
}

// NewProductRepository creates a new product repository instance
func NewProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return &productRepository{db: db}
}

// ListProducts retrieves paginated products ordered by created_at DESC
func (r *productRepository) ListProducts(page, pageSize int) ([]models.Product, error) {
	offset := (page - 1) * pageSize

	query := `
			WITH paged AS (
			  SELECT p.*
			  FROM products p
			  ORDER BY p.created_at DESC
			  LIMIT $1 OFFSET $2
			)
			SELECT
			  p.id,
			  p.code,
			  p.name,
			  p.description,
			  p.created_at,
			  pt.id         AS "product_type.id",
			  pt.code       AS "product_type.code",
			  pt.name       AS "product_type.name",
			  pt.created_at AS "product_type.created_at",
			  (
				SELECT COALESCE(
				  jsonb_agg(
					jsonb_build_object(
					  'id',         c.id,
					  'code',       c.code,
					  'name',       c.name,
					  'hex',        c.hex,
					  'created_at', c.created_at
					)
					ORDER BY c.name
				  ) FILTER (WHERE c.id IS NOT NULL),
				  '[]'::jsonb
				)
				FROM products_colors pc
				JOIN colors c ON c.id = pc.color_id
				WHERE pc.product_id = p.id
			  ) AS colors
			
			FROM paged p
			JOIN product_types pt ON pt.id = p.product_type_id
			ORDER BY p.created_at DESC;
			`

	var products []models.Product
	if err := r.db.Select(&products, query, pageSize, offset); err != nil {
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
