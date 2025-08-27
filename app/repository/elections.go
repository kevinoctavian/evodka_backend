package repository

import (
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type ElectionRepository interface {
	Create(election *model.Election) error
}

type ElectionRepo struct {
	db *database.DB
}

// Create implements ElectionRepository.
func (e *ElectionRepo) Create(election *model.Election) error {
	id, err := utils.NewULID()
	if err != nil {
		return err
	}

	query := `INSERT INTO elections (id, school_id, start_at, end_at) VALUES ($1, $2, $3, $4)`
	_, err = e.db.Exec(query, id.String(), election.SchoolID, election.StartAt, election.EndAt)
	return err
}

func NewElectionRepo(db *database.DB) ElectionRepository {
	return &ElectionRepo{db}
}
