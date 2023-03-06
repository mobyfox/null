package null

import (
	"database/sql"
	"encoding/json"
)

// Byte is a custom type that embeds sql.NullByte.
type Byte struct {
	sql.NullByte
}

// MarshalJSON uses the Byte custom type and implements the json.Marshaler interface.
func (b Byte) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Byte)
	}

	return []byte(`null`), nil
}
