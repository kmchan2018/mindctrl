package errors

import (
	"fmt"
)

type ArgumentError struct {
	message string
}

func NewArgumentError(format string, a ...interface{}) *ArgumentError {
	return &ArgumentError{
		message: fmt.Sprintf(format, a...),
	}
}

func NewExcessArgumentError() *ArgumentError {
	return &ArgumentError{
		message: fmt.Sprintf("excess argument found"),
	}
}

func NewMissingArgumentError(name string) *ArgumentError {
	return &ArgumentError{
		message: fmt.Sprintf("argument %s not found", name),
	}
}

func NewInvalidArgumentError(name string, reason string) *ArgumentError {
	return &ArgumentError{
		message: fmt.Sprintf("argument %s invalid: %s", name, reason),
	}
}

func (err *ArgumentError) Error() string {
	return err.message
}

func (err *ArgumentError) Is(target error) bool {
	_, ok := target.(*ArgumentError)
	return ok
}
