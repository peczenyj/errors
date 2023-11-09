package errors_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/peczenyj/errors"
)

func TestWrapFunction(t *testing.T) {
	t.Parallel()

	t.Run("errors.Wrap", func(t *testing.T) {
		t.Parallel()

		testWrapFunction(t, errors.Wrap)
	})

	t.Run("errors.WithMessage alias", func(t *testing.T) {
		t.Parallel()

		testWrapFunction(t, errors.WithMessage)
	})
}

func TestWrapfFunction(t *testing.T) {
	t.Parallel()

	t.Run("errors.Wrapf", func(t *testing.T) {
		t.Parallel()

		testWrapfFunction(t, errors.Wrapf)
	})

	t.Run("errors.WithMessagef alias", func(t *testing.T) {
		t.Parallel()

		testWrapfFunction(t, errors.WithMessagef)
	})
}

func testWrapFunction(t *testing.T, wrap func(error, string) error) {
	t.Helper()

	t.Run("whould wrap a non nil error", func(t *testing.T) {
		t.Parallel()

		err := errors.New("foo")
		err = wrap(err, "bar")

		require.EqualError(t, err, "bar: foo")
	})

	t.Run("wrap nil should return nil", func(t *testing.T) {
		t.Parallel()

		err := wrap(nil, "message")
		require.NoError(t, err)
	})
}

func testWrapfFunction(t *testing.T, wrapf func(error, string, ...any) error) {
	t.Helper()

	t.Run("whould wrap a non nil error", func(t *testing.T) {
		t.Parallel()

		err := errors.New("foo")
		err = wrapf(err, "bar %d", 1)

		require.EqualError(t, err, "bar 1: foo")
	})

	t.Run("wrap nil should return nil", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrapf(nil, "message %d", 1)
		require.NoError(t, err)
	})
}

func TestCauseFunction(t *testing.T) {
	t.Parallel()

	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")

	require.NoError(t, errors.Cause(nil))

	require.Equal(t, e1, errors.Cause(e3))
	require.Equal(t, e1, errors.Cause(e2))
	require.Equal(t, e1, errors.Cause(e1))
}

var _ error = (*causerError)(nil)

type causerError struct {
	err error
}

func (c *causerError) Error() string {
	return "caused by: " + c.err.Error()
}

func (c *causerError) Cause() error {
	return c.err
}

func TestCauseFunction_causer(t *testing.T) {
	t.Parallel()

	e1 := errors.New("error")
	e2 := &causerError{err: e1}

	require.Equal(t, e1, errors.Cause(e2))
}
