package domain

import "clean-arch-hicoll/pkg/dto"

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	Address  string
}

type UserRepository interface {
	ValidateUsernameAvailable(username string) error
	ValidateEmailAvailable(email string) error
	AddNewUser(user User) error
	GetUserById(userId int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserById(user User) error
	DeleteUserById(userId int) error
}

type UserUsecase interface {
	AddNewUser(req dto.UserReqDTO) error
	GetUserById(userId int) (dto.UserResDTO, error)
	GetAllUsers() ([]dto.UserResDTO, error)
	UpdateUserById(userId int, req dto.UserReqDTO) error
	DeleteUserById(userId int) error
}
