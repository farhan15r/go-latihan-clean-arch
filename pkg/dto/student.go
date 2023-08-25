package dto

import (
	"clean-arch-hicoll/shared/exception"

	validation "github.com/go-ozzo/ozzo-validation"
)

type StudentDTO struct {
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

func (s StudentDTO) Validation() error {
	err := validation.ValidateStruct(&s,
		validation.Field(&s.Fullname, validation.Required),
		validation.Field(&s.Address),
		validation.Field(&s.Batch, validation.Required),
		validation.Field(&s.Birthdate, validation.Required),
		validation.Field(&s.Class, validation.Required),
		validation.Field(&s.SchoolName, validation.Required))

	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}
	return nil
}
