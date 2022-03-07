//Unit.test, запуск: go test testify_test.go
package reflection

import (
	"github.com/stretchr/testify/require"
	"testing"
	//reflection "go_level_2/go_level_2/lesson7/reflection"
	reflection "go_level_2/go_level_2/lesson7/reflection"
)

// Use testify
func TestReflection(t *testing.T) {

	valuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	var x = reflection.Reflection(valuesMap, &reflection.Structure)
	var y error = nil
	require.Equal(t, x, y, "The two errors should be the same.")
}
