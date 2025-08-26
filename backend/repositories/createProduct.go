package repositories

import (
	"context"
	"database/sql"
	"sort"

	"github.com/AmirAziziDev/product-management-system/models"
	repoif "github.com/AmirAziziDev/product-management-system/repositories/interfaces"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *productRepository) CreateProduct(ctx context.Context, p models.Product, colorIDs []int) (id int, err error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Validate FK: product_type exists
	ok, err := productTypeExists(ctx, tx, p.ProductType.ID)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, repoif.ErrProductTypeNotFound
	}

	// Validate colors exist (if provided)
	if len(colorIDs) > 0 {
		missing, err := missingColorIDs(ctx, tx, colorIDs)
		if err != nil {
			return 0, err
		}
		if len(missing) > 0 {
			sort.Ints(missing)
			return 0, repoif.ErrColorsNotFound
		}
	}

	// Insert product, return id
	if err = tx.QueryRowxContext(ctx, `
		INSERT INTO products (code, name, description, product_type_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, p.Code, p.Name, p.Description, p.ProductType.ID).Scan(&id); err != nil {
		// UNIQUE, FK, CHECK violations bubble up; handler maps pq.Error (e.g., 23505)
		return 0, err
	}

	// Attach colors (if any)
	if len(colorIDs) > 0 {
		if _, err = tx.ExecContext(ctx, `
			INSERT INTO products_colors (product_id, color_id)
			SELECT $1, x FROM unnest($2::int[]) AS t(x)
			ON CONFLICT DO NOTHING
		`, id, pq.Array(colorIDs)); err != nil {
			return 0, err
		}
	}

	// Commit tx
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func productTypeExists(ctx context.Context, q sqlx.ExtContext, id int) (bool, error) {
	var ok bool
	err := sqlx.GetContext(ctx, q, &ok, `SELECT EXISTS(SELECT 1 FROM product_types WHERE id=$1)`, id)
	return ok, err
}

func missingColorIDs(ctx context.Context, q sqlx.ExtContext, ids []int) ([]int, error) {
	if len(ids) == 0 {
		return []int{}, nil
	}

	var missing []int
	query := `
		WITH input AS (
			SELECT x::int AS id, ord
			FROM unnest($1::int[]) WITH ORDINALITY AS t(x, ord)
		)
		SELECT i.id
		FROM input i
		LEFT JOIN colors c ON c.id = i.id
		WHERE c.id IS NULL
		GROUP BY i.id
		ORDER BY MIN(i.ord);
	`

	if err := sqlx.SelectContext(ctx, q, &missing, query, pq.Array(ids)); err != nil {
		return nil, err
	}
	return missing, nil
}
