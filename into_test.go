package errors_test

import (
	"testing"

	"github.com/peczenyj/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntoFunction(t *testing.T) {
	t.Parallel()

	t.Run("nil error should return nil, false", func(t *testing.T) {
		t.Parallel()

		var err error
		terr, ok := errors.Into[TimeoutError](err)
		require.False(t, ok)
		require.NoError(t, terr)
	})

	t.Run("normal error should return nil, false", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")

		terr, ok := errors.Into[TimeoutError](err)
		require.False(t, ok)
		require.NoError(t, terr)
	})

	t.Run("timeout error should return itself", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")
		err = &timeoutError{err}

		terr, ok := errors.Into[TimeoutError](err)
		require.True(t, ok)
		require.EqualError(t, terr, "timeout: ops")
		assert.True(t, terr.Timeout())
	})

	t.Run("wrapped timeout error should return itself", func(t *testing.T) {
		t.Parallel()

		err := errors.New("ops")
		err = &timeoutError{err}
		err = errors.Wrap(err, "unexpected")

		terr, ok := errors.Into[TimeoutError](err)
		require.True(t, ok)
		require.EqualError(t, terr, "timeout: ops")
		assert.True(t, terr.Timeout())
	})
}
