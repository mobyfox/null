package null

import (
	"database/sql"
	"encoding/json"
)

// String is a custom type that embeds sql.NullString.
type String struct {
	sql.NullString
}

// MarshalJSON uses the String custom type and implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return []byte(`null`), nil
}
