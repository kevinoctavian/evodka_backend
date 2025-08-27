package repository

import (
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type VoteRepository interface {
	Create(vote *model.Vote) error
}

type VoteRepo struct {
	db *database.DB
}

// Create implements VoteRepository.
func (e *VoteRepo) Create(vote *model.Vote) error {
	id, err := utils.NewULID()
	if err != nil {
		return err
	}

	query := `INSERT INTO votes (id, school_id, user_id, candidate_id, school_id) VALUES ($1, $2, $3, $4, $5)`
	_, err = e.db.Exec(query, id.String())
	return err
}

func NewVoteRepo(db *database.DB) VoteRepository {
	return &VoteRepo{db}
}
