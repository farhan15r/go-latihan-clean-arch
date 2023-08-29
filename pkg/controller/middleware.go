package controller

import (
	"clean-arch-hicoll/config"
	"clean-arch-hicoll/pkg/usecase"
	"clean-arch-hicoll/shared/exception"
	"clean-arch-hicoll/shared/response"
	"clean-arch-hicoll/shared/token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	conf *config.Configuration
}

func NewMiddleware(conf *config.Configuration) *Middleware {
	return &Middleware{
		conf: conf,
	}
}

func (m Middleware) RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	tm := token.NewTokenManager(m.conf)
	rlu := usecase.NewRequireLoginUsecase(tm)

	return func(c echo.Context) error {
		value := c.Request().Header.Get("Authorization")

		if value == "" {
			return response.SetResponseError(c, 401, "unauthorized", "unauthorized")
		}

		splitToken := strings.Split(value, "Bearer ")
		if len(splitToken) != 2 {
			return response.SetResponseError(c, 401, "unauthorized", "unauthorized")
		}

		accessToken := splitToken[1]

		userId, err := rlu.ValidateAccessToken(accessToken)
		if err != nil {
			return err
		}

		c.Set("userId", userId)

		return next(c)
	}
}

func (m Middleware) NotFoundHandler(c echo.Context) error {
	return response.SetResponseError(c, 404, "not found", "not found")
}

func (m Middleware) ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if clientError(c, err) {
				return nil
			} else {
				internalServerError(c, err)
			}
		}
		return nil
	}
}

func clientError(c echo.Context, err error) bool {
	exception, ok := err.(exception.ClientError)
	if ok {
		response.SetResponseError(c, exception.Code, exception.Message, exception.Errors)

		return true
	} else {
		return false
	}
}

func internalServerError(c echo.Context, err error) {
	response.SetResponseError(c, http.StatusInternalServerError, "internal server error", "internal server error")
}
