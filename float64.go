package null

import (
	"database/sql"
	"encoding/json"
)

// Float64 is a custom type that embeds sql.NullFloat64
type Float64 struct {
	sql.NullFloat64
}

// MarshalJSON uses the Float64 custom type and implements the json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	}

	return []byte(`null`), nil
}
