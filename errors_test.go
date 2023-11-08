package errors_test

import (
	"io"
	"testing"

	"github.com/peczenyj/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	{
		err := errors.New("foo")
		err = errors.Errorf("bar: %w", err)

		require.EqualError(t, err, "bar: foo")
	}

	{
		err := errors.New("foo")
		err = errors.Errorf("bar: %w", err)
		err = errors.Errorf("baz: %w", err)

		require.EqualError(t, err, "baz: bar: foo")
	}
}

func TestIsFunction(t *testing.T) {
	t.Parallel()

	err := errors.Errorf("failure: %w", io.EOF)

	assert.True(t, errors.Is(io.EOF, io.EOF))
	assert.True(t, errors.Is(err, io.EOF))
	assert.False(t, errors.Is(err, errors.ErrUnsupported))
}

var (
	_ error = (*customErr)(nil)
	_ error = (*anotherCustomErr)(nil)
)

type customErr struct{}

func (*customErr) Error() string { return "custom error" }

func (*customErr) Timeout() bool { return true }

type anotherCustomErr struct{}

func (*anotherCustomErr) Error() string { return "another custom error" }

func TestAsFunction(t *testing.T) {
	t.Parallel()

	err := &customErr{}

	{
		var cerr *customErr

		require.True(t, errors.As(err, &cerr))
		assert.NotNil(t, cerr)
	}

	{
		var terr interface{ Timeout() bool }

		require.True(t, errors.As(err, &terr))
		require.NotNil(t, terr)
		assert.True(t, terr.Timeout())
	}

	{
		var cerr *anotherCustomErr

		require.False(t, errors.As(err, &cerr))
		assert.Nil(t, cerr)
	}
}

func TestAsFunction_wrappingError(t *testing.T) {
	t.Parallel()

	err := errors.Errorf("failure: %w", &customErr{})

	{
		var cerr *customErr

		require.True(t, errors.As(err, &cerr))
		assert.NotNil(t, cerr)
	}

	{
		var terr interface{ Timeout() bool }

		require.True(t, errors.As(err, &terr))
		require.NotNil(t, terr)
		assert.True(t, terr.Timeout())
	}

	{
		var cerr *anotherCustomErr

		require.False(t, errors.As(err, &cerr))
		assert.Nil(t, cerr)
	}
}

func TestJoinFunction(t *testing.T) {
	err1 := errors.New("foo")
	err2 := errors.New("bar")
	err3 := error(nil)

	err := errors.Join(err1, err2, err3)

	require.EqualError(t, err, "foo\nbar")
}
