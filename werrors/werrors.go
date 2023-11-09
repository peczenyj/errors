// werrors is a minimalistic package that providers only
// Wrap, Wrapf, WithMessage, WithMessagef, Cause and Into functions.
//
// This allow use the standard "errors" package for New / Errorf, Is, As, etc.
package werrors

import "github.com/peczenyj/errors"

// Wrap returns an error annotating err, it is the equivalent to
//
//	fmt.Errorf("%s: %w", message, err)
//
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf returns an formatted error annotating err, it is the equivalent to
//
//	message := fmt.Sprintf(format, args...)
//	fmt.Errorf("%s: %w", message, err)
//
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}

// WithMessage is an alias to Wrap.
func WithMessage(err error, message string) error {
	return errors.WithMessage(err, message)
}

// WithMessagef is an alias to Wrapf.
func WithMessagef(err error, format string, args ...any) error {
	return errors.WithMessagef(err, format, args...)
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
	return errors.Cause(err) //nolint: wrapcheck
}

// Into finds the first error in err's chain that matches target type T, and if so, returns it.
//
// Into is type-safe alternative to As.
func Into[T error](err error) (val T, ok bool) { //nolint: ireturn
	return errors.Into[T](err)
}
