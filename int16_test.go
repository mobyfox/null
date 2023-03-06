package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	int16Test     = `{"int16":32767}`
	int16NullTest = `{"int16":null}`
)

func TestInt16(t *testing.T) {
	t.Run("should return int16", func(t *testing.T) {
		type testStruct struct {
			Int16 Int16 `json:"int16"`
		}

		test := &testStruct{}
		test.Int16.Int16 = 32767
		test.Int16.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int16Test, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Int16 Int16 `json:"int16"`
		}

		test := &testStruct{}
		test.Int16.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, int16NullTest, string(result))
	})
}
