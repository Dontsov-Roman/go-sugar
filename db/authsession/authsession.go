package authsession

import (
	. "go-sugar/db"
)

// Auth struct
type Auth struct {
	UserID    int      `json:"UserID"`
	DeviceID  string   `json:"DeviceID"`
	Token     string   `json:"Token"`
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
}

// GetByDeviceID delegate to Repository
func GetByDeviceID(id string) (*Auth, error) {
	return Repo.GetByDeviceID(id)
}

// Save - create new AUTH session
func (a *Auth) Save() (*Auth, error) {
	_, err := Repo.CleanBeforeCreate(a)
	if err != nil {
		return nil, err
	}
	return Repo.Create(a)
}
