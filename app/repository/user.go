package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type UserRepository interface {
	Create(user *model.CreateUser) error
	All(limit uint, offset int) ([]*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Exists(email string) (bool, error)
	Update(id uuid.UUID, user *model.UpdateUser) error
	Delete(id uuid.UUID) error
}

type UserRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

// Create implements UserRepository.
func (r *UserRepo) Create(user *model.CreateUser) error {
	query := `INSERT INTO users (name, email, password_hash, role) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role)
	return err
}

// All implements UserRepository.
func (r *UserRepo) All(limit uint, offset int) ([]*model.User, error) {
	var users []*model.User
	query := `SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC`
	var err error

	if limit > 0 && offset >= 0 {
		query = fmt.Sprintf("%s LIMIT $1 OFFSET $2", query)
		err = r.db.Select(&users, query, limit, offset)
	} else {
		err = r.db.Select(&users, query)
	}

	return users, err
}

func (r *UserRepo) Exists(email string) (bool, error) {
	query := `SELECT exists (SELECT 1 from users WHERE email = $1)`

	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, err
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	query := `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`
	var user model.User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.PublicID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other error
	}
	return &user, nil
}

// Delete implements UserRepository.
func (r *UserRepo) Delete(id uuid.UUID) error {
	panic("unimplemented")
}

// FindByID implements UserRepository.
func (r *UserRepo) FindByID(id uuid.UUID) (*model.User, error) {
	uuidString := id.String()
	query := `SELECT * FROM users WHERE public_id = $1 AND deleted_at IS NULL`
	var user model.User
	err := r.db.QueryRow(query, uuidString).Scan(&user.ID, &user.PublicID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other error
	}
	return &user, nil
}

// Update implements UserRepository.
func (r *UserRepo) Update(id uuid.UUID, user *model.UpdateUser) error {
	panic("unimplemented")
}
