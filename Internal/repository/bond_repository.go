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
	order string,
	search string,
) ([]models.Bond, error) {

	var bonds []models.Bond

	var query string
	var err error

	if search != "" {

		query = fmt.Sprintf(`
	SELECT *
	FROM bonds
	WHERE
		similarity(name, $1) > 0.1
		OR similarity(issuer, $1) > 0.1
	ORDER BY %s %s
	LIMIT $2
	OFFSET $3
`, orderBy, order)

		err = r.db.Select(
			&bonds,
			query,
			search,
			limit,
			offset,
		)

	} else {

		query = fmt.Sprintf(`
		SELECT *
		FROM bonds
		ORDER BY %s %s
		LIMIT $1
		OFFSET $2
	`, orderBy, order)

		err = r.db.Select(
			&bonds,
			query,
			limit,
			offset,
		)
	}

	if err != nil {
		return nil, err
	}

	return bonds, nil
}

func (r *BondRepository) BondExists(
	id string,
) (bool, error) {

	var count int

	err := r.db.Get(
		&count,
		`
		SELECT COUNT(*)
		FROM bonds
		WHERE b_id = $1
		`,
		id,
	)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *BondRepository) GetBondCount() (int, error) {

	var count int

	err := r.db.Get(
		&count,
		`
		SELECT COUNT(*)
		FROM bonds
		`,
	)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *BondRepository) GetTotalBondCount() (
	int,
	error,
) {

	var count int

	err := r.db.Get(
		&count,
		`
		SELECT COUNT(*)
		FROM bonds
		`,
	)

	if err != nil {
		return 0, err
	}

	return count, nil
}
