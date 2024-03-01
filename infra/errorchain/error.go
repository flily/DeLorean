package errorchain

import (
	"fmt"
)

// Error is a base error type that can be used to derive custom error types
type Error struct {
	Message    string
	InnerError error
}

// NewError creates a new ErrorBase with the given message
func NewError(format string, args ...any) *Error {
	e := &Error{
		Message:    fmt.Sprintf(format, args...),
		InnerError: nil,
	}
	return e
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.InnerError
}

func (e *Error) Derive(format string, args ...any) *Error {
	err := &Error{
		Message:    fmt.Sprintf(format, args...),
		InnerError: e,
	}

	return err
}

// Derive creates a new ErrorBase with the given message and the given error as the inner error
func Derive(err error, format string, args ...any) *Error {
	if err == nil {
		return nil
	}

	return &Error{
		Message:    fmt.Sprintf(format, args...),
		InnerError: err,
	}
}
