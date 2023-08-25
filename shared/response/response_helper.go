package response

import "github.com/labstack/echo/v4"

type JsonReponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JsonReponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func SetResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, JsonReponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func SetResponseError(c echo.Context, statusCode int, message string, errors interface{}) error {
	return c.JSON(statusCode, JsonReponseError{
		Code:    statusCode,
		Message: message,
		Errors:  errors,
	})
}
