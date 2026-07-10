package services

import (
	"errors"
	"fmt"
	"math"
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
	search string,
) ([]models.Bond, error) {

	if page < 1 {
		return nil, errors.New(
			"page does not exist.",
		)
	}

	if limit < 1 {
		limit = 10
	}

	total, err := s.repo.GetTotalBondCount()

	if err != nil {
		return nil, err
	}

	totalPages := int(
		math.Ceil(
			float64(total) / float64(limit),
		),
	)

	if total > 0 && page > totalPages {
		return nil, errors.New(
			fmt.Sprintf(
				"page %d does not exist. last available page is %d",
				page,
				totalPages,
			),
		)
	}

	offset := (page - 1) * limit

	ratingOrder := `
		CASE rating
    		WHEN 'Sovereign' THEN 1
    		WHEN 'AAA' THEN 2
			WHEN 'AA+' THEN 3
			WHEN 'AA' THEN 4
			WHEN 'AA-' THEN 5
			WHEN 'A+' THEN 6
			WHEN 'A' THEN 7
			WHEN 'A-' THEN 8
			WHEN 'BBB+' THEN 9
			WHEN 'BBB' THEN 10
			WHEN 'BBB-' THEN 11
			WHEN 'BB+' THEN 12
			WHEN 'BB' THEN 13
			WHEN 'BB-' THEN 14
			WHEN 'B+' THEN 15
			WHEN 'B' THEN 16
			WHEN 'B-' THEN 17
			ELSE 999
		END
		`
	frequencyOrder := `
		CASE rating
    		WHEN 'Monthly' THEN 1
			WHEN 'Quarterly' THEN 2
			WHEN 'Half-Yearly' THEN 3
			WHEN 'Yearly' THEN 4
			ELSE 999
		END
		`

	sortMap := map[string]string{
		"name":           "name",
		"yield":          "yield",
		"frequency":      frequencyOrder,
		"rating":         ratingOrder,
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
		search,
	)
}

func (s *BondService) GetBondCount() (
	int,
	error,
) {

	return s.repo.GetBondCount()
}
