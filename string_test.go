package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stringTest     = `{"string":"hello, world"}`
	stringNullTest = `{"string":null}`
)

func TestString(t *testing.T) {
	t.Run("should return string", func(t *testing.T) {
		type testStruct struct {
			String String `json:"string"`
		}

		test := &testStruct{}
		test.String.String = "hello, world"
		test.String.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, stringTest, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			String String `json:"string"`
		}

		test := &testStruct{}
		test.String.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, stringNullTest, string(result))
	})
}
