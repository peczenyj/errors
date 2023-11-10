//go:build go1.21
// +build go1.21

package errors

import "errors"

// ErrUnsupported indicates that a requested operation cannot be performed,
// because it is unsupported.
// requires minimum go 1.21.
var ErrUnsupported = errors.ErrUnsupported
