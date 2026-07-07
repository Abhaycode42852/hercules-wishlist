package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"wishlist-backend/internal/models"
)

type WishlistRepository struct {
	db *sqlx.DB
}

func NewWishlistRepository(
	db *sqlx.DB,
) *WishlistRepository {
	return &WishlistRepository{
		db: db,
	}
}

func (r *WishlistRepository) GetWishlistCount() (int, error) {

	var count int

	err := r.db.Get(
		&count,
		"SELECT COUNT(*) FROM wishlists",
	)

	return count, err
}

func (r *WishlistRepository) CreateWishlist(
	name string,
) (*models.Wishlist, error) {

	var wishlist models.Wishlist

	query := `
		INSERT INTO wishlists(name)
		VALUES($1)
		RETURNING w_id, name, size
	`

	err := r.db.Get(
		&wishlist,
		query,
		name,
	)

	if err != nil {
		return nil, err
	}

	return &wishlist, nil
}

func (r *WishlistRepository) GetAllWishlists() ([]models.Wishlist, error) {

	var wishlists []models.Wishlist

	err := r.db.Select(
		&wishlists,
		`SELECT w_id, name, size FROM wishlists`,
	)

	if err != nil {
		return nil, err
	}

	return wishlists, nil
}

func (r *WishlistRepository) GetWishlistByID(
	id string,
) (*models.Wishlist, error) {

	var wishlist models.Wishlist

	err := r.db.Get(
		&wishlist,
		`
		SELECT
			w_id,
			name,
			size
		FROM wishlists
		WHERE w_id = $1
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &wishlist, nil
}

func (r *WishlistRepository) UpdateWishlist(
	id string,
	name string,
) error {

	_, err := r.db.Exec(
		`
		UPDATE wishlists
		SET name = $1
		WHERE w_id = $2
		`,
		name,
		id,
	)

	return err
}

func (r *WishlistRepository) WishlistNameExists(
	name string,
) (bool, error) {

	var count int

	err := r.db.Get(
		&count,
		`
		SELECT COUNT(*)
		FROM wishlists
		WHERE LOWER(name) = LOWER($1)
		`,
		name,
	)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *WishlistRepository) DeleteWishlist(
	id string,
) error {

	result, err := r.db.Exec(
		`
		DELETE FROM wishlists
		WHERE w_id = $1
		`,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("wishlist not found")
	}

	return nil
}
