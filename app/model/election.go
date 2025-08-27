package model

import "time"

type Election struct {
	ID        string    `db:"id"`
	SchoolID  string    `db:"school_id"`
	StartAt   time.Time `db:"start_at"`
	EndAt     time.Time `db:"end_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
}
