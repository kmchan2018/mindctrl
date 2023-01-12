package errors

import (
	"fmt"
)

type ExecutionError struct {
	message string
	wrapped error
}

func NewExecutionError(format string, a ...interface{}) *ExecutionError {
	return &ExecutionError{
		message: fmt.Sprintf(format, a...),
		wrapped: nil,
	}
}

func WrapExecutionError(err error, format string, a ...interface{}) *ExecutionError {
	return &ExecutionError{
		message: fmt.Sprintf(format, a...),
		wrapped: err,
	}
}

func (err *ExecutionError) Error() string {
	return err.message
}

func (err *ExecutionError) Unwrap() error {
	return err.wrapped
}

func (err *ExecutionError) Is(target error) bool {
	_, ok := target.(*ExecutionError)
	return ok
}
