package shared

import (
	"github.com/jmoiron/sqlx"
)

// InitializeProductsSchema creates the products table and indexes
func InitializeProductsSchema(db *sqlx.DB) error {
	schema := `
	CREATE TABLE products (
		id SERIAL PRIMARY KEY,
		code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
		name TEXT NOT NULL UNIQUE,
		description TEXT NULL,
		created_at TIMESTAMPTZ DEFAULT now()
	);

	CREATE INDEX idx_products_code ON products(code);
	CREATE INDEX idx_products_created_at ON products(created_at DESC);
	`

	_, err := db.Exec(schema)
	return err
}
