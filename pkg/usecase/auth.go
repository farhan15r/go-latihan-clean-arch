package usecase

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/dto"
	"clean-arch-hicoll/shared/exception"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	tm domain.TokenManager
	ar domain.AuthRepository
	ur domain.UserRepository
}

func NewAuthUsecase(
	tm domain.TokenManager,
	ar domain.AuthRepository,
	ur domain.UserRepository,
) domain.AuthUsecase {
	return &AuthUsecase{
		tm: tm,
		ar: ar,
		ur: ur,
	}
}

func (au *AuthUsecase) Login(req dto.AuthDTO) (domain.Auth, error) {
	err := req.Validate()

	auth := domain.Auth{}

	if err != nil {
		return auth, err
	}

	user, err := au.ur.GetUserByUsername(req.Username)
	if err != nil {
		return auth, exception.NewClientError("username or password is not valid", "username or password is not valid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return auth, exception.NewClientError("username or password is not valid", "username or password is not valid")
	}

	accessToken, err := au.tm.GenerateAccessToken(user.Id)
	if err != nil {
		return auth, err
	}

	refreshToken, err := au.tm.GenerateRefreshToken(user.Id)
	if err != nil {
		return auth, err
	}

	err = au.ar.AddRefreshToken(refreshToken)
	if err != nil {
		return auth, err
	}

	auth.AccessToken = accessToken
	auth.RefreshToken = refreshToken

	return auth, nil
}

func (au *AuthUsecase) Refresh(req dto.AuthRefreshReqDTO) (dto.AuthRefreshResDTO, error) {
	err := req.Validate()

	res := dto.AuthRefreshResDTO{}

	err = au.ar.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return res, err
	}

	userId, err := au.tm.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		au.ar.RemoveRefreshToken(req.RefreshToken)
		return res, err
	}

	accessToken, err := au.tm.GenerateAccessToken(userId)
	if err != nil {
		return res, err
	}

	res.AccessToken = accessToken
	return res, nil
}

func (au *AuthUsecase) Logut(req dto.AuthLogoutReqDTO) error {
	err := req.Validate()
	if err != nil {
		return err
	}

	err = au.ar.RemoveRefreshToken(req.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}
