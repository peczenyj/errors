package errors_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/peczenyj/errors"
)

func TestNewCtor(t *testing.T) {
	t.Parallel()

	err := errors.New("foo")

	require.EqualError(t, err, "foo")
}

func TestErrorfCtor(t *testing.T) {
	t.Parallel()

	err := errors.Errorf("foo %d", 1)

	require.EqualError(t, err, "foo 1")
}

func TestErrorfCtor_wrappingError(t *testing.T) {
	t.Parallel()

	t.Run("should wrap once", func(t *testing.T) {
		t.Parallel()

		err := errors.New("foo")
		err = errors.Errorf("bar: %w", err)

		require.EqualError(t, err, "bar: foo")
	})

	t.Run("should wrap twice", func(t *testing.T) {
		t.Parallel()

		err := errors.New("foo")
		err = errors.Errorf("bar: %w", err)
		err = errors.Errorf("baz: %w", err)

		require.EqualError(t, err, "baz: bar: foo")
	})
}

func TestIsFunction(t *testing.T) {
	t.Parallel()

	err := errors.Errorf("failure: %w", io.EOF)

	assert.True(t, errors.Is(io.EOF, io.EOF))
	assert.True(t, errors.Is(err, io.EOF))
	assert.False(t, errors.Is(err, errors.ErrUnsupported))
}

var (
	_ error = (*customError)(nil)
	_ error = (*anotherCustomError)(nil)
)

type customError struct{}

func (*customError) Error() string { return "custom error" }

func (*customError) Timeout() bool { return true }

type anotherCustomError struct{}

func (*anotherCustomError) Error() string { return "another custom error" }

func TestAsFunction(t *testing.T) {
	t.Parallel()

	t.Run("custom error should return itself", func(t *testing.T) {
		t.Parallel()

		err := &customError{}

		var cerr *customError

		require.True(t, errors.As(err, &cerr))
		assert.NotNil(t, cerr)
	})

	t.Run("custom error should return timeout error", func(t *testing.T) {
		t.Parallel()

		err := &customError{}

		var terr interface{ Timeout() bool }

		require.True(t, errors.As(err, &terr))
		require.NotNil(t, terr)
		assert.True(t, terr.Timeout())
	})

	t.Run("custom error should not return another custom error", func(t *testing.T) {
		t.Parallel()

		err := &customError{}

		var cerr *anotherCustomError

		require.False(t, errors.As(err, &cerr))
		assert.Nil(t, cerr)
	})
}

func TestAsFunction_wrappingError(t *testing.T) {
	t.Parallel()

	t.Run("custom error should convert to itself", func(t *testing.T) {
		t.Parallel()

		err := errors.Errorf("failure: %w", &customError{})

		var cerr *customError

		require.True(t, errors.As(err, &cerr))
		assert.NotNil(t, cerr)
	})

	t.Run("custom error should convert as timeout error", func(t *testing.T) {
		t.Parallel()

		err := errors.Errorf("failure: %w", &customError{})

		var terr interface{ Timeout() bool }

		require.True(t, errors.As(err, &terr))
		require.NotNil(t, terr)
		assert.True(t, terr.Timeout())
	})

	t.Run("custom error should not return another custom error", func(t *testing.T) {
		t.Parallel()

		err := errors.Errorf("failure: %w", &customError{})

		var cerr *anotherCustomError

		require.False(t, errors.As(err, &cerr))
		assert.Nil(t, cerr)
	})
}

func TestJoinFunction(t *testing.T) {
	t.Parallel()

	err1 := errors.New("foo")
	err2 := errors.New("bar")
	err3 := error(nil)

	err := errors.Join(err1, err2, err3)

	require.EqualError(t, err, "foo\nbar")
}
