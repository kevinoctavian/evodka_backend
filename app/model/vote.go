package model

import "time"

type Vote struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	CandidateID string    `db:"candidate_id"`
	SchoolID    string    `db:"school_id"`
	CreatedAt   time.Time `db:"created_at"`
}
