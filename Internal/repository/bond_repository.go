package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"wishlist-backend/internal/models"
)

type BondRepository struct {
	db *sqlx.DB
}

func NewBondRepository(db *sqlx.DB) *BondRepository {
	return &BondRepository{
		db: db,
	}
}

func (r *BondRepository) GetAllBonds(
	limit int,
	offset int,
	orderBy string,
) ([]models.Bond, error) {

	var bonds []models.Bond

	query := fmt.Sprintf(`
		SELECT *
		FROM bonds
		ORDER BY %s
		LIMIT $1
		OFFSET $2
	`, orderBy)

	err := r.db.Select(
		&bonds,
		query,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	return bonds, nil
}
