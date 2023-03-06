package null

import (
	"database/sql"
	"encoding/json"
)

// Int32 is a custom type that embeds sql.NullInt32.
type Int32 struct {
	sql.NullInt32
}

// MarshalJSON uses the Int32 custom type and implements the json.Marshaler interface.
func (i Int32) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int32)
	}

	return []byte(`null`), nil
}
