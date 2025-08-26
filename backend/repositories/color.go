package repositories

import (
	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/jmoiron/sqlx"
)

type ColorRepository interface {
	GetColors() ([]models.Color, error)
}

type colorRepository struct {
	db *sqlx.DB
}

func NewColorRepository(db *sqlx.DB) ColorRepository {
	return &colorRepository{db: db}
}

func (r *colorRepository) GetColors() ([]models.Color, error) {
	query := `
		SELECT 
		    id,
		    code,
		    name,
		    hex,
		    created_at
		FROM colors
		ORDER BY created_at DESC 
	`

	var colors []models.Color
	if err := r.db.Select(&colors, query); err != nil {
		return nil, err
	}

	return colors, nil
}
