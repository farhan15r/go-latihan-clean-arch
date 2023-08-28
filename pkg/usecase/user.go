package usecase

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/dto"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: ur,
	}
}

func (uu *UserUsecase) AddNewUser(req dto.UserReqDTO) error {
	err := req.Validation()
	if err != nil {
		return err
	}

	user := domain.User{}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	mapstructure.Decode(req, &user)
	user.Password = string(hasedPassword)

	err = uu.userRepository.ValidateUsernameAvailable(user.Username)
	if err != nil {
		return err
	}
	err = uu.userRepository.ValidateEmailAvailable(user.Email)
	if err != nil {
		return err
	}
	err = uu.userRepository.AddNewUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) GetUserById(userId int) (dto.UserResDTO, error) {
	user, err := uu.userRepository.GetUserById(userId)

	userRes := dto.UserResDTO{}
	if err != nil {
		return userRes, err
	}

	mapstructure.Decode(user, &userRes)

	return userRes, nil
}

func (uu *UserUsecase) GetAllUsers() ([]dto.UserResDTO, error) {
	users, err := uu.userRepository.GetAllUsers()

	usersRes := []dto.UserResDTO{}

	if err != nil {
		return usersRes, err
	}

	for _, user := range users {
		userRes := dto.UserResDTO{}
		mapstructure.Decode(user, &userRes)
		usersRes = append(usersRes, userRes)
	}

	return usersRes, nil
}

func (uu *UserUsecase) UpdateUserById(userId int, req dto.UserReqDTO) error {
	err := req.Validation()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{}
	mapstructure.Decode(req, &user)
	user.Id = userId
	user.Password = string(hashedPassword)

	userOld, err := uu.userRepository.GetUserById(user.Id)
	if err != nil {
		return err
	}

	if user.Username != userOld.Username {
		err = uu.userRepository.ValidateUsernameAvailable(user.Username)
		if err != nil {
			return err
		}
	}

	if user.Email != userOld.Email {
		err = uu.userRepository.ValidateEmailAvailable(user.Email)
		if err != nil {
			return err
		}
	}

	err = uu.userRepository.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) DeleteUserById(userId int) error {
	_, err := uu.userRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	err = uu.userRepository.DeleteUserById(userId)
	if err != nil {
		return err
	}
	return nil
}
