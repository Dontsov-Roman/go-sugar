package authenticate

import (
	. "../../db"
)

// Auth struct
type Auth struct {
	UserID    int      `json:"UserID"`
	DeviceID  string   `json: "DeviceID`
	Token     string   `json:"Token"`
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
}
