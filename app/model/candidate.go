package model

import "time"

type Candidate struct {
	ID         string    `db:"id"`
	SchoolID   string    `db:"public_id"`
	ElectionID string    `db:"election_id"`
	KetuaName  string    `db:"ketua_name"`
	WakilName  string    `db:"wakil_name"`
	PhotoUrl   string    `db:"photo_url"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
