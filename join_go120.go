//go:build go1.20
// +build go1.20

package errors

import "errors"

// Join returns an error that wraps the given errors.
// requires minimum go 1.20.
func Join(errs ...error) error {
	return errors.Join(errs...)
}
