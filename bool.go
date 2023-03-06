package null

import (
	"database/sql"
	"encoding/json"
)

// Bool is a custom type that embeds sql.NullBool.
type Bool struct {
	sql.NullBool
}

// MarshalJSON uses the Bool custom type and implements the json.Marshaler interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}

	return []byte(`null`), nil
}
