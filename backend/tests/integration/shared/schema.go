package shared

import (
	"github.com/jmoiron/sqlx"
)

// InitializeProductsSchema creates the products table and indexes
func InitializeProductsSchema(db *sqlx.DB) error {
	schema := `
	-- Create product_types table
	CREATE TABLE product_types (
		id SERIAL PRIMARY KEY,
		code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
		name TEXT NULL,
		created_at TIMESTAMPTZ DEFAULT now()
	);

	-- Create colors table
	CREATE TABLE colors (
		id SERIAL PRIMARY KEY,
		code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
		name TEXT NOT NULL,
		hex TEXT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT now()
	);

	-- Create products table with product_type_id foreign key
	CREATE TABLE products (
		id SERIAL PRIMARY KEY,
		code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
		name TEXT NOT NULL UNIQUE,
		description TEXT NULL,
		product_type_id INTEGER NOT NULL REFERENCES product_types(id),
		created_at TIMESTAMPTZ DEFAULT now()
	);

	-- Create products_colors junction table
	CREATE TABLE products_colors (
		id SERIAL PRIMARY KEY,
		product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
		color_id INTEGER NOT NULL REFERENCES colors(id) ON DELETE CASCADE,
		created_at TIMESTAMPTZ DEFAULT now(),
		UNIQUE(product_id, color_id)
	);

	-- Create indexes
	CREATE INDEX idx_products_code ON products(code);
	CREATE INDEX idx_products_created_at ON products(created_at DESC);
	CREATE INDEX idx_product_types_code ON product_types(code);
	CREATE INDEX idx_colors_code ON colors(code);
	CREATE INDEX idx_products_colors_product_id ON products_colors(product_id);
	CREATE INDEX idx_products_colors_color_id ON products_colors(color_id);
	`

	_, err := db.Exec(schema)
	return err
}
