package dto

import (
	"clean-arch-hicoll/shared/exception"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AuthDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a AuthDTO) Validate() error {
	err := validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Password, validation.Required))

	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}
	return nil
}

type AuthRefreshReqDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (a AuthRefreshReqDTO) Validate() error {
	err := validation.ValidateStruct(&a,
		validation.Field(&a.RefreshToken, validation.Required))

	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}
	return nil
}

type AuthRefreshResDTO struct {
	AccessToken string `json:"access_token"`
}

type AuthLogoutReqDTO = AuthRefreshReqDTO
