package usecase

import "clean-arch-hicoll/pkg/domain"

type RequireLoginUsecase struct {
	tm domain.TokenManager
}

func NewRequireLoginUsecase(tm domain.TokenManager) domain.RequireLoginUsecase {
	return &RequireLoginUsecase{
		tm: tm,
	}
}

func (rlu *RequireLoginUsecase) ValidateAccessToken(token string) (int, error) {
	userId, err := rlu.tm.ValidateAccessToken(token)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
