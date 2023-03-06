package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	byteTest     = `{"byte":0}`
	byteNullTest = `{"byte":null}`
)

func TestByte(t *testing.T) {
	t.Run("should return byte", func(t *testing.T) {
		type testStruct struct {
			Byte Byte `json:"byte"`
		}

		test := &testStruct{}
		test.Byte.Byte = byte(0)
		test.Byte.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, byteTest, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Byte Byte `json:"byte"`
		}

		test := &testStruct{}
		test.Byte.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, byteNullTest, string(result))
	})
}
