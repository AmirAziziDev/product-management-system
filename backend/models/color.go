package models

import "time"

type Color struct {
	ID        int       `json:"id" db:"id"`
	Code      int       `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Hex       string    `json:"hex" db:"hex"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
