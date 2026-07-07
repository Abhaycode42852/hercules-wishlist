package services

import (
	"errors"

	"wishlist-backend/internal/models"
	"wishlist-backend/internal/repository"
)

type BondService struct {
	repo *repository.BondRepository
}

func NewBondService(
	repo *repository.BondRepository,
) *BondService {
	return &BondService{
		repo: repo,
	}
}

func (s *BondService) GetAllBonds(
	page int,
	limit int,
	sort string,
	order string,
) ([]models.Bond, error) {

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	sortMap := map[string]string{
		"name":           "name",
		"yield":          "yield",
		"frequency":      "frequency",
		"rating":         "rating",
		"min_units":      "min_units",
		"min_investment": "min_investment",
		"tenure": `
			(EXTRACT(YEAR FROM maturity_date)
			- EXTRACT(YEAR FROM CURRENT_DATE))
		`,
	}

	orderBy, exists := sortMap[sort]

	if !exists {
		if sort != "" {
			return nil, errors.New("invalid sort parameter")
		}

		orderBy = "name"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	return s.repo.GetAllBonds(
		limit,
		offset,
		orderBy,
		order,
	)
}

func (s *BondService) GetBondCount() (
	int,
	error,
) {

	return s.repo.GetBondCount()
}
