package db

import (
	"database/sql"
	"encoding/json"

	"github.com/go-sql-driver/mysql"
)

// NullTime new simple NullTime
type NullTime struct {
	mysql.NullTime
}

// NullString new simple NullString
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullTime
func (ni *NullTime) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Time)
}

// MarshalJSON for NullTime
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}
