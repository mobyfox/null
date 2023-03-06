package null

import (
	"database/sql"
	"encoding/json"
)

// Int16 is a custom type that embeds sql.NullInt16.
type Int16 struct {
	sql.NullInt16
}

// MarshalJSON uses the Int16 custom type and implements the json.Marshaler interface.
func (i Int16) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int16)
	}

	return []byte(`null`), nil
}
