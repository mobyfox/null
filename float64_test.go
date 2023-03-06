package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	float64Test     = `{"float":3.14159265359}`
	float64NullTest = `{"float":null}`
)

func TestFloat64(t *testing.T) {
	t.Run("should return float", func(t *testing.T) {
		type testStruct struct {
			Float Float64 `json:"float"`
		}

		test := &testStruct{}
		test.Float.Float64 = 3.14159265359
		test.Float.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, float64Test, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Float Float64 `json:"float"`
		}

		test := &testStruct{}
		test.Float.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, float64NullTest, string(result))
	})
}
