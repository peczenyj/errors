package errors_test

import (
	"testing"

	"github.com/peczenyj/errors"
	"github.com/stretchr/testify/require"
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

	{
		err := errors.New("foo")
		err = wrap(err, "bar")

		require.EqualError(t, err, "bar: foo")
	}

	{
		err := wrap(nil, "message")
		require.NoError(t, err)
	}
}

func testWrapfFunction(t *testing.T, wrapf func(error, string, ...any) error) {
	t.Helper()
	{
		err := errors.New("foo")
		err = errors.Wrapf(err, "bar %d", 1)

		require.EqualError(t, err, "bar 1: foo")
	}

	{
		err := errors.Wrapf(nil, "message %d", 1)
		require.NoError(t, err)
	}
}

func TestCauseFunction(t *testing.T) {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")

	require.NoError(t, errors.Cause(nil))

	require.Equal(t, e1, errors.Cause(e3))
	require.Equal(t, e1, errors.Cause(e2))
	require.Equal(t, e1, errors.Cause(e1))
}

var _ error = (*causerErr)(nil)

type causerErr struct {
	err error
}

func (c *causerErr) Error() string {
	return "caused by: " + c.err.Error()
}

func (c *causerErr) Cause() error {
	return c.err
}

func TestCauseFunction_causer(t *testing.T) {
	e1 := errors.New("error")
	e2 := &causerErr{err: e1}

	require.Equal(t, e1, errors.Cause(e2))
}
