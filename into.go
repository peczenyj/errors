package errors

import "errors"

// Into finds the first error in err's chain that matches target type T, and if so, returns it.
//
// Into is type-safe alternative to As.
//
// this function was ported from
// https://github.com/go-faster/errors
func Into[T error](err error) (val T, ok bool) { //nolint: ireturn
	ok = errors.As(err, &val)

	return val, ok
}
