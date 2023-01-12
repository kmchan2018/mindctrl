package errors

type FlagError struct {
	wrapped error
}

func NewFlagError(err error) *FlagError {
	return &FlagError{
		wrapped: err,
	}
}

func (err *FlagError) Error() string {
	return err.wrapped.Error()
}

func (err *FlagError) Unwrap() error {
	return err.wrapped
}

func (err *FlagError) Is(target error) bool {
	_, ok := target.(*FlagError)
	return ok
}
