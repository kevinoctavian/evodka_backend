package repository

import (
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type CandidateRepository interface {
	Create(candidate *model.Candidate) error
}

type CandidateRepo struct {
	db *database.DB
}

// Create implements CandidateRepository.
func (e *CandidateRepo) Create(candidate *model.Candidate) error {
	id, err := utils.NewULID()
	if err != nil {
		return err
	}

	query := `INSERT INTO candidates (id, school_id, elections_id, ketua_name, wakil_name, photo_url) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = e.db.Exec(query, id.String(), candidate.SchoolID, candidate.ElectionID, candidate.KetuaName, candidate.WakilName, candidate.PhotoUrl)
	return err
}

func NewCandidateRepo(db *database.DB) CandidateRepository {
	return &CandidateRepo{db}
}
