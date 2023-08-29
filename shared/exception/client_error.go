package exception

import "strconv"

type ClientError struct {
	Code    int
	Status  string
	Message string
	Errors  interface{}
}

func (clientError ClientError) Error() string {
	return strconv.Itoa(clientError.Code) + clientError.Message
}

func NewNotFoundError(message string, errors interface{}) ClientError {
	return ClientError{
		Code:    404,
		Status:  "NOT FOUND",
		Message: message,
		Errors:  errors,
	}
}

func NewClientError(message string, errors interface{}) ClientError {
	return ClientError{
		Code:    400,
		Status:  "BAD REQUEST",
		Message: message,
		Errors:  errors,
	}
}

func NewUnauthorizedError(message string, errors interface{}) ClientError {
	return ClientError{
		Code:    401,
		Status:  "UNAUTHORIZED",
		Message: message,
		Errors:  errors,
	}
}

func NewForbiddenError(message string, errors interface{}) ClientError {
	return ClientError{
		Code:    403,
		Status:  "FORBIDDEN",
		Message: message,
		Errors:  errors,
	}
}
