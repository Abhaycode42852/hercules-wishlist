package services

import (
	"errors"
	"strings"
	"wishlist-backend/internal/models"
	"wishlist-backend/internal/repository"
)

type WishlistService struct {
	repo     *repository.WishlistRepository
	bondRepo *repository.BondRepository
}

func NewWishlistService(
	repo *repository.WishlistRepository,
	bondRepo *repository.BondRepository,
) *WishlistService {
	return &WishlistService{
		repo:     repo,
		bondRepo: bondRepo,
	}
}

func (s *WishlistService) CreateWishlist(
	name string,
) (*models.Wishlist, error) {

	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New(
			"wishlist name cannot be empty",
		)
	}

	if len(name) > 30 {
		return nil, errors.New(
			"wishlist name cannot exceed 30 characters",
		)
	}

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
) (*models.WishlistDetails, error) {

	return s.repo.GetWishlistWithBonds(id) //this feature needs renaming but works fine without it
}

func (s *WishlistService) UpdateWishlist(
	id string,
	name string,
) error {
	exists, err :=
		s.repo.WishlistNameExists(name)

	if err != nil {
		return err
	}

	if exists {
		return errors.New(
			"name already exists choose a diffrent name",
		)
	}

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

func (s *WishlistService) AddBondToWishlist(
	wishlistID string,
	bondID string,
) error {

	exists, err :=
		s.repo.WishlistExists(wishlistID)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New(
			"wishlist not found",
		)
	}

	bondExists, err :=
		s.bondRepo.BondExists(bondID)

	if err != nil {
		return err
	}

	if !bondExists {
		return errors.New(
			"bond not found",
		)
	}

	duplicate, err :=
		s.repo.BondAlreadyInWishlist(
			wishlistID,
			bondID,
		)

	if err != nil {
		return err
	}

	if duplicate {
		return errors.New(
			"bond already exists in wishlist",
		)
	}

	size, err :=
		s.repo.GetWishlistSize(
			wishlistID,
		)

	if err != nil {
		return err
	}

	if size >= 10 {
		return errors.New(
			"wishlist already contains maximum 10 bonds",
		)
	}

	return s.repo.AddBondToWishlist(
		wishlistID,
		bondID,
	)
}

func (s *WishlistService) RemoveBondFromWishlist(
	wishlistID string,
	bondID string,
) error {

	return s.repo.RemoveBondFromWishlist(
		wishlistID,
		bondID,
	)
}
