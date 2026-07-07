package models

type WishlistDetails struct {
	Wishlist

	Bonds []Bond `json:"bonds"`
}
