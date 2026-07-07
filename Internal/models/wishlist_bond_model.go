package models

import "github.com/google/uuid"

type WishlistBond struct {
	WID uuid.UUID `db:"w_id" json:"w_id"`

	BID uuid.UUID `db:"b_id" json:"b_id"`
}
