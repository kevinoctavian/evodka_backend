package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID         int64     `db:"id"`
	UserID     uuid.UUID `db:"user_public_id"`
	Token      string    `db:"token"` // hashed
	DeviceName string    `db:"device_name"`
	IPAddress  string    `db:"ip_address"`
	ExpiresAt  time.Time `db:"expires_at"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// id BIGSERIAL PRIMARY KEY,
// user_public_id UUID NOT NULL, -- untuk referensi publik
// token TEXT NOT NULL, -- refresh token yang di-hash (bukan raw token)
// device_name VARCHAR(100), -- nama device misalnya: "Chrome on Windows"
// ip_address VARCHAR(45),   -- IPv4/IPv6 address
// expires_at TIMESTAMPTZ NOT NULL,
// created_at TIMESTAMPTZ DEFAULT now(),
// updated_at TIMESTAMPTZ DEFAULT now(),
