package models

import (
	"time"

	"github.com/google/uuid"
)

type Bond struct {
	BID uuid.UUID `db:"b_id" json:"b_id"`

	Name string `db:"name" json:"name"`

	Yield float64 `db:"yield" json:"yield"`

	Frequency string `db:"frequency" json:"frequency"`

	Rating string `db:"rating" json:"rating"`

	MinUnits int `db:"min_units" json:"min_units"`

	MaxUnits int `db:"max_units" json:"max_units"`

	MaturityDate time.Time `db:"maturity_date" json:"maturity_date"`

	Issuer string `db:"issuer" json:"issuer"`

	ISIN string `db:"isin" json:"isin"`

	CouponRate float64 `db:"coupon_rate" json:"coupon_rate"`

	LogoURL string `db:"logo_url" json:"logo_url"`

	MinInvestment float64 `db:"min_investment" json:"min_investment"`

	Sector string `db:"sector" json:"sector"`

	FaceValue float64 `db:"face_value" json:"face_value"`
}
