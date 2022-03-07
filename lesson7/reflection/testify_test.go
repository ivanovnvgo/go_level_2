package reflection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Use testify
func TestReflection(t *testing.T) {

	type In struct {
		Key int
	}

	var Structure In

	valuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	var x = Reflection(valuesMap, Structure)
	var y error = nil
	require.Equal(t, x, y, "The two errors should be the same.")
}
