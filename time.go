package null

import (
	"database/sql"
	"encoding/json"
)

// Time is a custom type that embeds sql.NullTime.
type Time struct {
	sql.NullTime
}

// MarshalJSON uses the Time custom type and implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}

	return []byte(`null`), nil
}
