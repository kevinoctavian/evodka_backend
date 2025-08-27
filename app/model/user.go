package model

import (
	"time"
)

type User struct {
	ID           string    `db:"id"`
	SchoolId     string    `db:"school_id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

type CreateUser struct {
	SchoolId string `form:"school_id"`
	Username string `form:"username" validate:"required,lte=50,gte=5"`
	Email    string `form:"email" validate:"required,email,lte=150"`
	Password string `form:"password" validate:"required,lte=100,gte=8"`
	Role     string `form:"role"`
}

type LoginUser struct {
	Email    string `form:"email" validate:"required,email,lte=150"`
	Password string `form:"password" validate:"required,lte=100,gte=8"`
}

type UpdateUser struct {
	SchoolId string `form:"school_id"`
	Username string `form:"username" validate:"omitempty,lte=50,gte=5"`
	Password string `form:"password" validate:"omitempty,lte=100,gte=8"`
	Role     string `form:"role"`
}
