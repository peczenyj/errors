package errors

import "errors"

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}

// ErrUnsupported indicates that a requested operation cannot be performed,
// because it is unsupported.
var ErrUnsupported = errors.ErrUnsupported

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Is reports whether any error in err's tree matches target.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Join returns an error that wraps the given errors.
func Join(errs ...error) error {
	return errors.Join(errs...)
}
