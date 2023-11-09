package errors_test

import (
	"fmt"

	"github.com/peczenyj/errors"
)

func ExampleNew() {
	err := errors.New("whoops")
	fmt.Println(err)

	// Output: whoops
}

func ExampleWrap() {
	cause := errors.New("whoops")
	err := errors.Wrap(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func ExampleWithMessage() {
	cause := errors.New("whoops")
	err := errors.WithMessage(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")

	return errors.Wrap(e3, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(errors.Cause(err))

	// Output: outer: middle: inner: error
	// error
}

type timeoutError struct {
	err error
}

func (te *timeoutError) Error() string {
	return "timeout: " + te.err.Error()
}

func (te *timeoutError) Timeout() bool {
	return true
}

type TimeoutError interface {
	Timeout() bool
	error
}

func ExampleInto() {
	err := errors.New("ops")
	err = &timeoutError{err}
	err = errors.Wrap(err, "unexpected")

	if terr, ok := errors.Into[TimeoutError](err); ok {
		fmt.Println(terr)
	}

	// Output: timeout: ops
}
