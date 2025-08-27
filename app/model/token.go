package model

import (
	"time"
)

type RefreshToken struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	Token      string    `db:"token_hash"` // hashed
	DeviceName string    `db:"device_name"`
	IPAddress  string    `db:"ip_address"`
	UserAgent  string    `db:"user_agent"`
	ExpiresAt  time.Time `db:"expires_at"`
	CreatedAt  time.Time `db:"created_at"`
	Revoked    bool      `db:"revoked"`
}
