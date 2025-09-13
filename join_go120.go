//go:build go1.20

package errors

import "errors"

// Join returns an error that wraps the given errors.
func Join(errs ...error) error {
	return errors.Join(errs...)
}
