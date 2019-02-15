package authenticate

import (
	. "../../db"
)

// Auth struct
type Auth struct {
	UserID    int      `json:"UserID"`
	Token     string   `json:"Token"`
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
}
