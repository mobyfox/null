package null

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	timeTest     = `{"time":"2023-03-06T13:55:00Z"}`
	timeNullTest = `{"time":null}`
)

func TestTime(t *testing.T) {
	t.Run("should return time", func(t *testing.T) {
		type testStruct struct {
			Time Time `json:"time"`
		}

		test := &testStruct{}
		test.Time.Time = time.Date(2023, time.March, 6, 13, 55, 00, 0, time.UTC)
		test.Time.Valid = true

		result, _ := json.Marshal(test)
		assert.JSONEq(t, timeTest, string(result))
	})

	t.Run("should return null", func(t *testing.T) {
		type testStruct struct {
			Time Time `json:"time"`
		}

		test := &testStruct{}
		test.Time.Valid = false

		result, _ := json.Marshal(test)
		assert.JSONEq(t, timeNullTest, string(result))
	})
}
