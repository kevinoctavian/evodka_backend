package repository

import (
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type TokenRepository interface {
	Create(token *model.RefreshToken) error
	GetByPublicID(publicID string) (*model.RefreshToken, error)
	FindByToken(token string) (*model.RefreshToken, error)
	Update(token *model.RefreshToken) error
	Delete(id int64) error
	Exists(publicID string) (bool, error)
}

type TokenRepo struct {
	db *database.DB
}

// FindByToken implements TokenRepository.
func (t *TokenRepo) FindByToken(token string) (*model.RefreshToken, error) {
	query := `SELECT * FROM refresh_tokens WHERE token = $1`
	var refreshToken model.RefreshToken
	err := t.db.QueryRow(query, token).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.DeviceName,
		&refreshToken.IPAddress,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
		&refreshToken.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

// Create implements TokenRepository.
func (t *TokenRepo) Create(token *model.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (
	user_public_id, 
	token,
	device_name,
	ip_address,
	expires_at,
	created_at, 
	updated_at
	) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`

	if len(token.DeviceName) > 100 {
		token.DeviceName = token.DeviceName[:100] // Truncate to 100 characters
	}

	err := t.db.QueryRow(query, token.UserID, token.Token, token.DeviceName, token.IPAddress, token.ExpiresAt).Scan(&token.ID)
	if err != nil {
		return err
	}
	token.CreatedAt = token.UpdatedAt // Set CreatedAt to now since we just created it
	return nil
}

// Delete implements TokenRepository.
func (t *TokenRepo) Delete(id int64) error {
	panic("unimplemented")
}

// Exists implements TokenRepository.
func (t *TokenRepo) Exists(publicID string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM refresh_tokens WHERE public_id = $1)`
	var exists bool
	err := t.db.QueryRow(query, publicID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GetByPublicID implements TokenRepository.
func (t *TokenRepo) GetByPublicID(publicID string) (*model.RefreshToken, error) {
	panic("unimplemented")
}

// Update implements TokenRepository.
func (t *TokenRepo) Update(token *model.RefreshToken) error {
	query := `UPDATE refresh_tokens SET 
		token = $1, 
		device_name = $2, 
		ip_address = $3, 
		expires_at = $4, 
		updated_at = NOW() 
		WHERE id = $5`

	if len(token.DeviceName) > 100 {
		token.DeviceName = token.DeviceName[:100] // Truncate to 100 characters
	}

	_, err := t.db.Exec(query, token.Token, token.DeviceName, token.IPAddress, token.ExpiresAt, token.ID)
	return err
}

func NewTokenRepo(db *database.DB) TokenRepository {
	return &TokenRepo{db}
}
