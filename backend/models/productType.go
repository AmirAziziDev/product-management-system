package models

import "time"

type ProductType struct {
	ID        int       `json:"id" db:"id"`
	Code      int       `json:"code" db:"code"`
	Name      *string   `json:"name,omitempty" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
