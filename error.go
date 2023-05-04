package runtime

import (
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

var (
	TooManyError       = Error{Code: http.StatusTooManyRequests, Message: http.StatusText(http.StatusTooManyRequests)}
	UnauthorizedError  = Error{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)}
	AlreadyExistsError = Error{Code: http.StatusUnprocessableEntity, Message: http.StatusText(http.StatusFound)}
	NotFoundError      = Error{Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound)}
	InvalidError       = Error{Code: http.StatusUnprocessableEntity, Message: ""}
)

func Duplicate(msg string) Error {
	err := AlreadyExistsError
	err.Message = msg

	return err
}

func NotFound(msg string) Error {
	err := NotFoundError
	err.Message = msg

	return err
}

func Invalid(msg string) Error {
	err := InvalidError
	err.Message = msg

	return err
}
