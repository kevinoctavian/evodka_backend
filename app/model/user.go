package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           int64      `db:"id"`
	PublicID     uuid.UUID  `db:"public_id"`
	Name         string     `db:"name"`
	Email        string     `db:"email"`
	PasswordHash string     `db:"password_hash"`
	Role         string     `db:"role"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at,omitzero"`
}

func NewUser() *User {
	return &User{}
}

type CreateUser struct {
	Username string `json:"username" validate:"required,lte=50,gte=5"`
	Email    string `json:"email" validate:"required,email,lte=150"`
	Password string `json:"password" validate:"required,lte=100,gte=8"`
	Role     string `json:"role"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email,lte=150"`
	Password string `json:"password" validate:"required,lte=100,gte=8"`
}

type UpdateUser struct {
	Username string `json:"username" validate:"omitempty,lte=50,gte=5"`
	Password string `json:"password" validate:"omitempty,lte=100,gte=8"`
	Role     string `json:"role"`
}
