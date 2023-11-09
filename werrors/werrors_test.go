package werrors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/peczenyj/errors/werrors"
)

func TestWrapFunction(t *testing.T) {
	t.Parallel()

	t.Run("werrors.Wrap", func(t *testing.T) {
		t.Parallel()

		testWrapFunction(t, werrors.Wrap)
	})

	t.Run("werrors.WithMessage alias", func(t *testing.T) {
		t.Parallel()

		testWrapFunction(t, werrors.WithMessage)
	})
}

func TestWrapfFunction(t *testing.T) {
	t.Parallel()

	t.Run("werrors.Wrapf", func(t *testing.T) {
		t.Parallel()

		testWrapfFunction(t, werrors.Wrapf)
	})

	t.Run("werrors.WithMessagef alias", func(t *testing.T) {
		t.Parallel()

		testWrapfFunction(t, werrors.WithMessagef)
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

		err := werrors.Wrapf(nil, "message %d", 1)
		require.NoError(t, err)
	})
}

func TestCauseFunction(t *testing.T) {
	t.Parallel()

	e1 := errors.New("error")
	e2 := werrors.Wrap(e1, "inner")
	e3 := werrors.Wrap(e2, "middle")

	require.NoError(t, werrors.Cause(nil))

	require.Equal(t, e1, werrors.Cause(e3))
	require.Equal(t, e1, werrors.Cause(e2))
	require.Equal(t, e1, werrors.Cause(e1))
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

	require.Equal(t, e1, werrors.Cause(e2))
}

func TestIntoFunction(t *testing.T) {
	t.Parallel()

	t.Run("nil error should return nil, false", func(t *testing.T) {
		t.Parallel()

		var err error
		terr, ok := werrors.Into[TimeoutError](err)
		require.False(t, ok)
		require.NoError(t, terr)
	})

	t.Run("normal error should return nil, false", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")

		terr, ok := werrors.Into[TimeoutError](err)
		require.False(t, ok)
		require.NoError(t, terr)
	})

	t.Run("timeout error should return itself", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")
		err = &timeoutError{err}

		terr, ok := werrors.Into[TimeoutError](err)
		require.True(t, ok)
		require.EqualError(t, terr, "timeout: ops")
		assert.True(t, terr.Timeout())
	})

	t.Run("wrapped timeout error should return itself", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")
		err = &timeoutError{err}
		err = werrors.Wrap(err, "unexpected")

		terr, ok := werrors.Into[TimeoutError](err)
		require.True(t, ok)
		require.EqualError(t, terr, "timeout: ops")
		assert.True(t, terr.Timeout())
	})
}
