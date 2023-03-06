package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	int32Test     = `{"int32":2147483647}`
	int32NullTest = `{"int32":null}`
)

func TestInt32(t *testing.T) {
	t.Run("should return int32", func(t *testing.T) {
		type testStruct struct {
			Int32 Int32 `json:"int32"`
		}

		test := &testStruct{}
		test.Int32.Int32 = 2147483647
		test.Int32.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int32Test, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Int32 Int32 `json:"int32"`
		}

		test := &testStruct{}
		test.Int32.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int32NullTest, string(result))
	})
}
