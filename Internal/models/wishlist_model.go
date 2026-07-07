package models

import (
	"github.com/google/uuid"
)

type Wishlist struct {
	WID uuid.UUID `db:"w_id" json:"w_id"`

	Name string `db:"name" json:"name"`

	Size int `db:"size" json:"size"`
}
