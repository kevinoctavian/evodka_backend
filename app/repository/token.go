package repository

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/pkg/utils"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

type TokenRepository interface {
	Create(token *model.RefreshToken) error
	GetByPublicID(publicID string) (*model.RefreshToken, error)
	FindByToken(token string) (*model.RefreshToken, error)
	Update(token *model.RefreshToken) error
	DeleteByToken(token string) error
	Exists(publicID string) (bool, error)
}

type TokenRepo struct {
	db *database.DB
}

// FindByToken implements TokenRepository.
func (t *TokenRepo) FindByToken(token string) (*model.RefreshToken, error) {
	tokenHash := sha256.Sum256([]byte(token))
	tokenHashString := hex.EncodeToString(tokenHash[:])

	query := `SELECT * FROM refresh_tokens WHERE token_hash = $1`
	var refreshToken model.RefreshToken
	err := t.db.QueryRow(query, tokenHashString).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.DeviceName,
		&refreshToken.IPAddress,
		&refreshToken.UserAgent,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
		&refreshToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

// Create implements TokenRepository.
func (t *TokenRepo) Create(token *model.RefreshToken) error {
	id, err := utils.NewULID()
	if err != nil {
		return err
	}

	query := `INSERT INTO refresh_tokens (
		id,
		user_id, 
		token_hash,
		device_name,
		ip_address,
		user_agent,
		expires_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	if len(token.DeviceName) > 100 {
		token.DeviceName = token.DeviceName[:100] // Truncate to 100 characters
	}

	tokenHash := sha256.Sum256([]byte(token.Token))
	token.Token = hex.EncodeToString(tokenHash[:])

	err = t.db.QueryRow(
		query,
		id.String(),
		token.UserID,
		token.Token,
		token.DeviceName,
		token.IPAddress,
		token.UserAgent,
		token.ExpiresAt,
	).Scan(&token.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements TokenRepository.
func (t *TokenRepo) DeleteByToken(token string) error {
	query := `DELETE FROM refresh_tokens WHERE token_hash = $1`
	_, err := t.db.Exec(query, token)
	return err
}

// Exists implements TokenRepository.
func (t *TokenRepo) Exists(publicID string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM refresh_tokens WHERE user_id = $1)`
	var exists bool
	err := t.db.QueryRow(query, publicID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GetByPublicID implements TokenRepository.
func (t *TokenRepo) GetByPublicID(publicID string) (*model.RefreshToken, error) {
	query := `SELECT * FROM refresh_tokens WHERE user_id = $1`
	var refreshToken model.RefreshToken
	err := t.db.QueryRow(query, publicID).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.DeviceName,
		&refreshToken.IPAddress,
		&refreshToken.UserAgent,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
		&refreshToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

// Update implements TokenRepository.
func (t *TokenRepo) Update(token *model.RefreshToken) error {
	query := `UPDATE refresh_tokens SET 
		token_hash = $1,
		device_name = $2, 
		ip_address = $3,
		expires_at = $4
		WHERE user_id = $5`

	if len(token.DeviceName) > 100 {
		token.DeviceName = token.DeviceName[:100] // Truncate to 100 characters
	}

	tokenHash := sha256.Sum256([]byte(token.Token))
	token.Token = hex.EncodeToString(tokenHash[:])

	_, err := t.db.Exec(query, token.Token, token.DeviceName, token.IPAddress, token.ExpiresAt, token.UserID)
	return err
}

func NewTokenRepo(db *database.DB) TokenRepository {
	return &TokenRepo{db}
}
