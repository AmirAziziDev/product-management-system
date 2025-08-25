package models

import "time"

type Product struct {
	ID          int         `json:"id" db:"id"`
	Code        int         `json:"code" db:"code"`
	Name        string      `json:"name" db:"name"`
	Description *string     `json:"description,omitempty" db:"description"`
	ProductType ProductType `json:"product_type,omitempty" db:"product_type"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
}
