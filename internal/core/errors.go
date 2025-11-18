package core

import "errors"

type ErrorType string

const (
	ValidationError ErrorType = "VALIDATION_ERROR"
	AuthError       ErrorType = "AUTH_ERROR"
	NotFoundError   ErrorType = "NOT_FOUND"
	ServerError     ErrorType = "SERVER_ERROR"
	DomainError     ErrorType = "DOMAIN_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(t ErrorType, msg string) *AppError {
	return &AppError{
		Type:    t,
		Message: msg,
	}
}

func Is(err error, t ErrorType) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Type == t
}

func Wrap(err error, t ErrorType) *AppError {
	return &AppError{
		Type:    t,
		Message: err.Error(),
	}
}

var ErrUnauthorized = errors.New("unauthorized")
