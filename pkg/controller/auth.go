package controller

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/dto"
	"clean-arch-hicoll/shared/exception"
	"clean-arch-hicoll/shared/response"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
}

func (ac AuthController) PostAuth(c echo.Context) error {
	credentials := dto.AuthDTO{}

	err := c.Bind(&credentials)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	jwt, err := ac.AuthUsecase.Login(credentials)
	if err != nil {
		return err
	}

	return response.SetResponse(c, 200, "success", jwt)
}

func (ac AuthController) PutAuth(c echo.Context) error {
	req := dto.AuthRefreshReqDTO{}

	err := c.Bind(&req)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	res, err := ac.AuthUsecase.Refresh(req)
	if err != nil {
		return err
	}

	return response.SetResponse(c, 200, "success", res)
}

func (ac AuthController) DeleteAuth(c echo.Context) error {
	req := dto.AuthLogoutReqDTO{}

	err := c.Bind(&req)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	err = ac.AuthUsecase.Logut(req)
	if err != nil {
		return err
	}

	return response.SetResponse(c, 200, "success", nil)
}
