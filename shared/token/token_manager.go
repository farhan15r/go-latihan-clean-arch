package token

import (
	"clean-arch-hicoll/config"
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/exception"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	conf *config.Configuration
}

func NewTokenManager(config *config.Configuration) domain.TokenManager {
	return &TokenManager{
		conf: config,
	}
}

func (tm *TokenManager) GenerateAccessToken(userId int) (string, error) {
	exp := time.Now().Add(time.Minute * time.Duration(tm.conf.AccessTokenExp))

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(tm.conf.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}

func (tm *TokenManager) GenerateRefreshToken(userId int) (string, error) {
	exp := time.Now().Add(time.Minute * time.Duration(tm.conf.RefreshTokenExp))

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(tm.conf.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}

func (tm *TokenManager) ValidateAccessToken(token string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(tm.conf.AccessTokenSecret), nil
	})
	if err != nil {
		return 0, exception.NewUnauthorizedError("token is not valid", err.Error())
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return 0, exception.NewUnauthorizedError("access token is not valid", "access is not valid")
	}

	userId := claims["userId"].(float64)
	return int(userId), nil
}

func (tm *TokenManager) ValidateRefreshToken(token string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(tm.conf.RefreshTokenSecret), nil
	})
	if err != nil {
		return 0, exception.NewClientError("token is not valid", err.Error())
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return 0, exception.NewClientError("refesh token is not valid", "token is not valid")
	}

	userId := claims["userId"].(float64)
	return int(userId), nil
}
