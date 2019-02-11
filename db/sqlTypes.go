package db

import (
	"encoding/json"

	"github.com/go-sql-driver/mysql"
)

// NullTime new simple NullTime
type NullTime struct {
	mysql.NullTime
}

// MarshalJSON for NullTime
func (ni *NullTime) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Time)
}
