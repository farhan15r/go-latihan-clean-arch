package controller

import (
	"clean-arch-hicoll/shared/exception"
	"clean-arch-hicoll/shared/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
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
