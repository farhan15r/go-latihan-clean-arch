package dto

import (
	"clean-arch-hicoll/shared/exception"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserReqDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (user UserReqDTO) Validation() error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 0)),
		validation.Field(&user.Address),
	)

	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}
	return nil
}

type UserResDTO struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
