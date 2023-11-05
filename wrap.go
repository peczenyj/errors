package errors

import (
	"fmt"
)

// Wrap ...
func Wrap(err error, msg string) error {
	if err != nil {
		return wrap(err, msg)
	}

	return nil
}

// Wrapf ...
func Wrapf(err error, format string, args ...any) error {
	if err != nil {
		return wrap(err, fmt.Sprintf(format, args...))
	}

	return nil
}

func wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Cause ...
func Cause(err error) error {
	if err == nil {
		return nil
	}

	var causerErr causer

	if As(err, &causerErr) {
		return causerErr.Cause()
	}

	previous := err
	cause := Unwrap(err)
	for cause != nil {
		previous = cause
		cause = Unwrap(cause)
	}

	return previous
}

type causer interface {
	Cause() error
}
