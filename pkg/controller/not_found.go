package controller

import (
	"clean-arch-hicoll/shared/response"

	"github.com/labstack/echo/v4"
)

func NotFoundHandler(c echo.Context) error {
	return response.SetResponseError(c, 404, "not found", "not found")
}
