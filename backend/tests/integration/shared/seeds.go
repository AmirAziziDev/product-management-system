package shared

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// SeedProductData inserts test product data into the database
func SeedProductData(db *sqlx.DB) error {
	// Insert test products with some variety in timestamps
	products := []struct {
		code        int
		name        string
		description *string
	}{
		{101, "Bookcase", stringPtr("Perfect for organizing books and displaying decorative items")},
		{102, "Bed Frame High Oak", nil},
		{103, "Daybed Frame", stringPtr("Versatile seating and sleeping solution for small spaces")},
		{104, "Shelf Unit", nil},
		{105, "Storage Unit", stringPtr("Ideal for organizing household items and keeping spaces tidy")},
		{106, "Shelving Unit Pine", stringPtr("Natural pine wood construction with multiple storage compartments")},
		{107, "Coffee Table", nil},
		{108, "Sleeper Sectional", stringPtr("Comfortable seating that converts to a bed for guests")},
		{109, "Armchair Birch", stringPtr("Elegant single seat chair with birch wood frame")},
		{110, "Sectional Sofa", nil},
	}

	for _, product := range products {
		_, err := db.Exec(
			"INSERT INTO products (code, name, description) VALUES ($1, $2, $3)",
			product.code, product.name, product.description,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product %s: %w", product.name, err)
		}
	}

	return nil
}

// stringPtr is a helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
}
