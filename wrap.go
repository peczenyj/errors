package errors

import (
	"errors"
	"fmt"
)

// Wrap returns an error annotating err, it is the equivalent to
//
//	fmt.Errorf("%s: %w", message, err)
//
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err != nil {
		return wrap(err, message)
	}

	return nil
}

// Wrap returns an formatted error annotating err, it is the equivalent to
//
//	message := fmt.Sprintf(format, args...)
//	fmt.Errorf("%s: %w", message, err)
//
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...any) error {
	if err != nil {
		return wrap(err, fmt.Sprintf(format, args...))
	}

	return nil
}

// WithMessage is an alias to Wrap.
func WithMessage(err error, message string) error {
	return Wrap(err, message)
}

// WithMessagef is an alias to Wrapf.
func WithMessagef(err error, format string, args ...any) error {
	return Wrapf(err, format, args...)
}

func wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, we will call Unwrap in a loop until
// find the original error.
// If the error is nil, nil will be returned without further investigation.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	type causer interface {
		Cause() error
	}

	var causerErr causer

	if errors.As(err, &causerErr) {
		return causerErr.Cause() //nolint: wrapcheck
	}

	cause, next := err, Unwrap(err)
	for next != nil {
		cause, next = next, Unwrap(next)
	}

	return cause
}
