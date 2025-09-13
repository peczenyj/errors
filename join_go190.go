//go:build !go1.20

package errors

import (
	"fmt"
	"strings"
)

var (
	_ error     = joinError(nil)
	_ unwrapper = joinError(nil)
)

type joinError []error

func (je joinError) Error() string {
	if len(je) == 1 {
		return je[0].Error()
	}

	var buff strings.Builder

	for index, err := range je {
		if index > 0 {
			fmt.Fprintln(&buff)
		}

		fmt.Fprint(&buff, err.Error())
	}

	return buff.String()
}

func (je joinError) Unwrap() []error {
	return []error(je)
}

type unwrapper interface {
	Unwrap() []error
}

// Join returns an error that wraps the given errors.
func Join(errs ...error) error {
	notNilErrs := errs[:0]

	for _, err := range errs {
		if err != nil {
			notNilErrs = append(notNilErrs, err)
		}
	}

	if len(notNilErrs) == 0 {
		return nil
	}

	if len(notNilErrs) == 1 {
		if err, ok := notNilErrs[0].(unwrapper); ok {
			return err
		}
	}

	for i := len(notNilErrs); i < len(errs); i++ {
		errs[i] = nil
	}

	return joinError(notNilErrs)
}
