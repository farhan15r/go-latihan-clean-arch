package response

import "github.com/labstack/echo/v4"

type JsonReponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, JsonReponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}
