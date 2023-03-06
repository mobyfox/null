package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	boolTest     = `{"bool":true}`
	boolNullTest = `{"bool":null}`
)

func TestBool(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		type testStruct struct {
			Bool Bool `json:"bool"`
		}

		test := &testStruct{}
		test.Bool.Bool = true
		test.Bool.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, boolTest, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Bool Bool `json:"bool"`
		}

		test := &testStruct{}
		test.Bool.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, boolNullTest, string(result))
	})
}
