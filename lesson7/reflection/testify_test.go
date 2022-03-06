package reflection

//package fibonacci

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Use testify
func TestReflectionExample(t *testing.T) {

	type In struct {
		Key int
	}

	var structure In

	valuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	var x = structure.Reflection(valuesMap)
	var y error = nil
	require.Equal(t, x, y, "The two errors should be the same.")
}
