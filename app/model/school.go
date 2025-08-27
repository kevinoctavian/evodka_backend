package model

import "time"

type School struct {
	ID        string    `db:"id"`
	Name      string    `db:"name" form:"name"`
	Type      string    `db:"type" form:"type"`
	NPSN      string    `db:"npsn" form:"npsn"`
	Address   string    `db:"address" form:"address"`
	Logo      string    `db:"logo" form:"logo"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
