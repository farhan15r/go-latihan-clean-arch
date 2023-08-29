package repository

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/exception"
	"database/sql"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (ar *AuthRepository) AddRefreshToken(token string) error {
	query := "INSERT INTO authentications (refresh_token) VALUES ($1)"

	_, err := ar.DB.Exec(query, token)

	if err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) ValidateRefreshToken(token string) error {
	query := "SELECT refresh_token FROM authentications WHERE refresh_token = $1"

	row := ar.DB.QueryRow(query, token)

	var refreshToken string
	err := row.Scan(&refreshToken)

	if err != nil {
		return exception.NewClientError("refresh token is not valid", "refresh token is not valid")
	}

	return nil
}

func (ar *AuthRepository) RemoveRefreshToken(token string) error {
	query := "DELETE FROM authentications WHERE refresh_token = $1"

	_, err := ar.DB.Exec(query, token)

	if err != nil {
		return err
	}
	return nil
}
