package usecase

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/exception"
	"strings"
)

type RequireLoginUsecase struct {
	tm domain.TokenManager
}

func NewRequireLoginUsecase(tm domain.TokenManager) domain.RequireLoginUsecase {
	return &RequireLoginUsecase{
		tm: tm,
	}
}

func (rlu *RequireLoginUsecase) ValidateAccessToken(authHeader string) (int, error) {
	if authHeader == "" {
		return 0, exception.NewUnauthorizedError("unauthorized", "unauthorized")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return 0, exception.NewUnauthorizedError("unauthorized", "unauthorized")
	}

	accessToken := splitToken[1]

	userId, err := rlu.tm.ValidateAccessToken(accessToken)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
