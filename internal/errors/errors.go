package errors

import "errors"

const (
	ValidationError = "validation error"
	AlreadyExists   = "already exists"
	NotFound        = "not found"
)

func ErrorValidationError() error {
	return errors.New(ValidationError)
}

func IsValidationError(err error) bool {
	return err.Error() == ValidationError
}

func ErrorAlreadyExists() error {
	return errors.New(AlreadyExists)
}

func ErrorNotFound() error {
	return errors.New(NotFound)
}

func IsAlreadyExists(err error) bool {
	return err.Error() == AlreadyExists
}

func IsNotFound(err error) bool {
	return err.Error() == NotFound
}
