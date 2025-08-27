package repository

import (
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type SchoolRepository interface {
	Create(school *model.School) error
}

type SchoolRepo struct {
	db *database.DB
}

func NewSchoolRepo(db *database.DB) SchoolRepository {
	return &SchoolRepo{db}
}

func (r *SchoolRepo) Create(school *model.School) error {
	id, err := utils.NewULID()
	if err != nil {
		return err
	}

	query := `INSERT INTO schools (
		id,
		name,
		type,
		npsn,
		address,
		logo
	) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = r.db.Exec(query, id.String(), school.Name, school.Type, school.NPSN, school.Address, school.Logo)
	return err
}
