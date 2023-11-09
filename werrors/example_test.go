package werrors_test

import (
	"errors"
	"fmt"

	"github.com/peczenyj/errors/werrors"
)

func ExampleWrap() {
	cause := errors.New("whoops")
	err := werrors.Wrap(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func ExampleWithMessage() {
	cause := errors.New("whoops")
	err := werrors.WithMessage(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func fn() error {
	e1 := errors.New("error")
	e2 := werrors.Wrap(e1, "inner")
	e3 := werrors.Wrap(e2, "middle")

	return werrors.Wrap(e3, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(werrors.Cause(err))

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
	err = werrors.Wrap(err, "unexpected")

	if terr, ok := werrors.Into[TimeoutError](err); ok {
		fmt.Println(terr)
	}

	// Output: timeout: ops
}
