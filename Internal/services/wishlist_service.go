package services

import (
	"errors"

	"wishlist-backend/internal/models"
	"wishlist-backend/internal/repository"
)

type WishlistService struct {
	repo *repository.WishlistRepository
}

func NewWishlistService(
	repo *repository.WishlistRepository,
) *WishlistService {
	return &WishlistService{
		repo: repo,
	}
}

func (s *WishlistService) CreateWishlist(
	name string,
) (*models.Wishlist, error) {

	exists, err :=
		s.repo.WishlistNameExists(name)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New(
			"wishlist name already exists",
		)
	}

	count, err :=
		s.repo.GetWishlistCount()

	if err != nil {
		return nil, err
	}

	if count >= 5 {
		return nil, errors.New(
			"maximum 5 wishlists allowed",
		)
	}

	return s.repo.CreateWishlist(name)
}
func (s *WishlistService) GetAllWishlists() (
	[]models.Wishlist,
	error,
) {
	return s.repo.GetAllWishlists()
}

func (s *WishlistService) GetWishlistByID(
	id string,
) (*models.Wishlist, error) {

	return s.repo.GetWishlistByID(id)
}

func (s *WishlistService) UpdateWishlist(
	id string,
	name string,
) error {

	return s.repo.UpdateWishlist(
		id,
		name,
	)
}

func (s *WishlistService) DeleteWishlist(
	id string,
) error {

	return s.repo.DeleteWishlist(id)
}
