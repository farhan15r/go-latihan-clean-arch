package domain

import "clean-arch-hicoll/pkg/dto"

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenManager interface {
	GenerateAccessToken(userId int) (string, error)
	GenerateRefreshToken(userId int) (string, error)
	ValidateAccessToken(token string) (int, error)
	ValidateRefreshToken(token string) (int, error)
}

type AuthRepository interface {
	AddRefreshToken(token string) error
	ValidateRefreshToken(token string) error
	RemoveRefreshToken(token string) error
}

type AuthUsecase interface {
	Login(req dto.AuthDTO) (Auth, error)
	Refresh(req dto.AuthRefreshReqDTO) (dto.AuthRefreshResDTO, error)
	Logut(req dto.AuthLogoutReqDTO) error
}
