package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Color struct {
	ID        int       `json:"id" db:"id"`
	Code      int       `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Hex       string    `json:"hex" db:"hex"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ColorList is a []Color that can scan JSON/JSONB from Postgres.
type ColorList []Color

// Scan implements sql.Scanner so sqlx can decode JSON/JSONB into []Color.
func (cl *ColorList) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, cl)
	case string:
		return json.Unmarshal([]byte(v), cl)
	case nil:
		*cl = ColorList{}
		return nil
	default:
		return fmt.Errorf("ColorList.Scan: unsupported type %T", v)
	}
}
