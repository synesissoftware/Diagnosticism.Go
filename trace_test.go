package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_EnableTracing(t *testing.T) {

	go func() {
		current := d.EnableTracing(false)
		defer d.EnableTracing(current)

		require.False(t, current, "`EnableTracing()` should return `false` when first called")
	}()

	go func() {
		current := d.EnableTracing(true)
		defer d.EnableTracing(current)

		require.False(t, current, "`EnableTracing()` should return `false` when first called")

		next := d.EnableTracing(false)

		require.True(t, next, "`EnableTracing()` should return `true` from the first call")
	}()
}
