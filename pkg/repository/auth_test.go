package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func truncateAuthenticationsTable() {
	_, err := dbConn.Exec("TRUNCATE TABLE authentications")
	if err != nil {
		panic(err)
	}
}

func TestAuthRepository_AddRefreshToken(t *testing.T) {
	defer truncateAuthenticationsTable()

	t.Run("should success add refresh token", func(t *testing.T) {
		repo := NewAuthRepository(dbConn)
		err := repo.AddRefreshToken("token")

		assert.NoError(t, err)

		result := ""
		dbConn.QueryRow("SELECT refresh_token FROM authentications").Scan(&result)

		assert.Equal(t, "token", result)
	})
}

func TestAuthRepository_ValidateRefreshToken(t *testing.T) {
	defer truncateAuthenticationsTable()

	t.Run("should error validate refresh token", func(t *testing.T) {
		repo := NewAuthRepository(dbConn)

		err := repo.ValidateRefreshToken("token")

		assert.Error(t, err)
	})

	t.Run("should success validate refresh token", func(t *testing.T) {
		repo := NewAuthRepository(dbConn)

		repo.AddRefreshToken("token")

		err := repo.ValidateRefreshToken("token")

		assert.NoError(t, err)
	})
}

func TestAuthRepository_RemoveRefreshToken(t *testing.T) {
	defer truncateAuthenticationsTable()

	t.Run("should success remove refresh token", func(t *testing.T) {
		repo := NewAuthRepository(dbConn)

		repo.AddRefreshToken("token")

		err := repo.RemoveRefreshToken("token")

		assert.NoError(t, err)

		result := ""
		dbConn.QueryRow("SELECT refresh_token FROM authentications").Scan(&result)

		assert.Equal(t, "", result)
	})
}
