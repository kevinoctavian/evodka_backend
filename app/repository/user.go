package repository

import (
	"database/sql"
	"fmt"

	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *model.CreateUser) error
	All(limit uint, offset int) ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Exists(email string) (bool, error)
	Update(id string, user *model.UpdateUser) error
	Delete(id string) error
}

type UserRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

// Create implements UserRepository.
func (r *UserRepo) Create(user *model.CreateUser) error {
	id, err := utils.NewULID()
	if err != nil {
		return nil
	}

	query := `INSERT INTO users (id, school_id, username, email, password_hash, role) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.db.Exec(query, id.String(), user.SchoolId, user.Username, user.Email, user.Password, user.Role)
	fmt.Println(user)
	return err
}

// All implements UserRepository.
func (r *UserRepo) All(limit uint, offset int) ([]*model.User, error) {
	var users []*model.User
	query := `SELECT * FROM users ORDER BY created_at DESC`
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
	query := `SELECT * FROM users WHERE email = $1`
	var user model.User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.SchoolId,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other error
	}
	return &user, nil
}

// Delete implements UserRepository.
func (r *UserRepo) Delete(id string) error {
	panic("unimplemented")
}

// FindByID implements UserRepository.
func (r *UserRepo) FindByID(id string) (*model.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.SchoolId, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other error
	}
	return &user, nil
}

// Update implements UserRepository.
func (r *UserRepo) Update(id string, user *model.UpdateUser) error {
	// Build dynamic query based on non-nil fields in user
	setClauses := []string{}
	args := []any{}
	argIdx := 1

	if user.Username != "" {
		setClauses = append(setClauses, fmt.Sprintf("username = $%d", argIdx))
		args = append(args, user.Username)
		argIdx++
	}
	if user.Password != "" {
		setClauses = append(setClauses, fmt.Sprintf("password_hash = $%d", argIdx))
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		args = append(args, hashedPassword)
		argIdx++
	}
	if user.Role != "" {
		setClauses = append(setClauses, fmt.Sprintf("role = $%d", argIdx))
		args = append(args, user.Role)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil // Nothing to update
	}

	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE id = $%d",
		fmt.Sprint(setClauses[0]),
		argIdx,
	)
	for i := 1; i < len(setClauses); i++ {
		query = fmt.Sprintf("%s, %s", query[:len("UPDATE users SET ")+len(fmt.Sprint(setClauses[0]))], setClauses[i]) + query[len("UPDATE users SET ")+len(fmt.Sprint(setClauses[0])):]
	}
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
