package shared

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// SeedProductData inserts test product data into the database
func SeedProductData(db *sqlx.DB) error {
	// Insert test product types
	productTypes := []struct {
		code int
		name *string
	}{
		{1, stringPtr("Furniture")},
		{2, stringPtr("Storage")},
		{3, stringPtr("Seating")},
		{4, stringPtr("Tables")},
	}

	for _, productType := range productTypes {
		_, err := db.Exec(
			"INSERT INTO product_types (code, name) VALUES ($1, $2)",
			productType.code, productType.name,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product type %v: %w", productType.name, err)
		}
	}

	// Insert test colors
	colors := []struct {
		code int
		name string
		hex  string
	}{
		{1, "White", "#FFFFFF"},
		{2, "Black", "#000000"},
		{3, "Brown", "#8B4513"},
		{4, "Oak", "#D2B48C"},
		{5, "Pine", "#FDF5E6"},
		{6, "Birch", "#F5F5DC"},
		{7, "Gray", "#808080"},
		{8, "Blue", "#0000FF"},
	}

	for _, color := range colors {
		_, err := db.Exec(
			"INSERT INTO colors (code, name, hex) VALUES ($1, $2, $3)",
			color.code, color.name, color.hex,
		)
		if err != nil {
			return fmt.Errorf("failed to insert color %s: %w", color.name, err)
		}
	}

	// Insert test products with product_type_id
	products := []struct {
		code          int
		name          string
		description   *string
		productTypeID int
	}{
		{101, "Bookcase", stringPtr("Perfect for organizing books and displaying decorative items"), 2}, // Storage
		{102, "Bed Frame High Oak", nil, 1}, // Furniture
		{103, "Daybed Frame", stringPtr("Versatile seating and sleeping solution for small spaces"), 3}, // Seating
		{104, "Shelf Unit", nil, 2}, // Storage
		{105, "Storage Unit", stringPtr("Ideal for organizing household items and keeping spaces tidy"), 2},            // Storage
		{106, "Shelving Unit Pine", stringPtr("Natural pine wood construction with multiple storage compartments"), 2}, // Storage
		{107, "Coffee Table", nil, 4}, // Tables
		{108, "Sleeper Sectional", stringPtr("Comfortable seating that converts to a bed for guests"), 3}, // Seating
		{109, "Armchair Birch", stringPtr("Elegant single seat chair with birch wood frame"), 3},          // Seating
		{110, "Sectional Sofa", nil, 3}, // Seating
	}

	for _, product := range products {
		_, err := db.Exec(
			"INSERT INTO products (code, name, description, product_type_id) VALUES ($1, $2, $3, $4)",
			product.code, product.name, product.description, product.productTypeID,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product %s: %w", product.name, err)
		}
	}

	// Insert product-color relationships
	productColors := []struct {
		productID int
		colorID   int
	}{
		{1, 1}, {1, 3}, // Bookcase: White, Brown
		{2, 4},         // Bed Frame High Oak: Oak
		{3, 1}, {3, 7}, // Daybed Frame: White, Gray
		{4, 1}, {4, 2}, // Shelf Unit: White, Black
		{5, 1}, {5, 7}, // Storage Unit: White, Gray
		{6, 5},         // Shelving Unit Pine: Pine
		{7, 3}, {7, 2}, // Coffee Table: Brown, Black
		{8, 7}, {8, 8}, // Sleeper Sectional: Gray, Blue
		{9, 6},           // Armchair Birch: Birch
		{10, 7}, {10, 2}, // Sectional Sofa: Gray, Black
	}

	for _, pc := range productColors {
		_, err := db.Exec(
			"INSERT INTO products_colors (product_id, color_id) VALUES ($1, $2)",
			pc.productID, pc.colorID,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product-color relationship product_id=%d, color_id=%d: %w", pc.productID, pc.colorID, err)
		}
	}

	return nil
}

// stringPtr is a helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
}
