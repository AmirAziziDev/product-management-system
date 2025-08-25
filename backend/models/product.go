package models

import "time"

type Product struct {
	ID          int       `json:"id" db:"id"`
	Code        int       `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
