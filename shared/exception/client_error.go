package exception

import "strconv"

type ClientError struct {
	Code    int
	Status  string
	Message string
}

func (clientError ClientError) Error() string {
	return strconv.Itoa(clientError.Code) + clientError.Message
}

func NewNotFoundError(message string) ClientError {
	return ClientError{
		Code:    404,
		Status:  "NOT FOUND",
		Message: message,
	}
}

func NewClientError(message string) ClientError {
	return ClientError{
		Code:    400,
		Status:  "BAD REQUEST",
		Message: message,
	}
}
