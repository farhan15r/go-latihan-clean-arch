package controller

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/dto"
	"clean-arch-hicoll/shared/exception"
	"clean-arch-hicoll/shared/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc UserController) PostUser(c echo.Context) error {
	user := dto.UserReqDTO{}

	err := c.Bind(&user)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	err = uc.UserUsecase.AddNewUser(user)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusCreated, "success", nil)
}

func (uc UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	user, err := uc.UserUsecase.GetUserById(id)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", user)
}

func (uc UserController) GetUsers(c echo.Context) error {
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", users)
}

func (uc UserController) PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	user := dto.UserReqDTO{}

	err = c.Bind(&user)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	err = uc.UserUsecase.UpdateUserById(id, user)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", nil)
}

func (uc UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	err = uc.UserUsecase.DeleteUserById(id)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", nil)
}
