package errors_test

import (
	"testing"

	"github.com/peczenyj/errors"

	"github.com/stretchr/testify/require"
)

func TestJoinFunction(t *testing.T) {
	t.Parallel()

	err1 := errors.New("foo")
	err2 := errors.New("bar")
	err3 := error(nil)

	err := errors.Join(err1, err2, err3)

	require.EqualError(t, err, "foo\nbar")
}
