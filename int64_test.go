package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	int64Test     = `{"int64":9223372036854775807}`
	int64NullTest = `{"int64":null}`
)

func TestInt64(t *testing.T) {
	t.Run("should return int64", func(t *testing.T) {
		type testStruct struct {
			Int64 Int64 `json:"int64"`
		}

		test := &testStruct{}
		test.Int64.Int64 = 9223372036854775807
		test.Int64.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int64Test, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Int64 Int64 `json:"int64"`
		}

		test := &testStruct{}
		test.Int64.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int64NullTest, string(result))
	})
}
