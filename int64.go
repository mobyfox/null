package null

import (
	"database/sql"
	"encoding/json"
)

// Int64 is a custom type that embeds sql.NullInt64.
type Int64 struct {
	sql.NullInt64
}

// MarshalJSON uses the Int64 custom type and implements the json.Marshaler interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	}

	return []byte(`null`), nil
}
