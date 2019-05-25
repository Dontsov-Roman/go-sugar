package users

import (
	"database/sql"
	"fmt"
	"strconv"

	. "../../db"
)

// User main struct
type User struct {
	ID        int      `json:"ID"`
	Name      string   `json:"Name"`
	Password  string   `json:"Password"`
	Type      int      `json:"Type"`
	Status    int      `json:"Status"`
	Email     string   `json:"Email"`
	Phone     string   `json:"Phone"`
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
	DeletedAt NullTime `json:"DeletedAt"`
}

// RegistrateByEmailUser get by first registrate via email
type RegistrateByEmailUser struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Status   int    `json:"Status"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
	DeviceID string `json:"DeviceID"`
}

// Save entity
func (u *User) Save() (*User, error) {
	if u.ID != 0 {
		ok, err := Repo.Update(u)
		if !ok {
			return nil, err
		}
		return u, err
	}
	return Repo.Create(u)
}

// Validate delegate to Repo
func (u *User) Validate() (bool, ValidateError) {
	return Repo.Validate(u)
}

// Delete entity
func (u *User) Delete() bool {
	return Repo.DeleteByID(strconv.Itoa(u.ID))
}

// CheckPassword checking password
func (u *User) CheckPassword(password string) bool {
	return Repo.CreateHash(password) == u.Password
}
func parseRows(rows *sql.Rows) []User {
	var users []User
	for rows.Next() {
		p, err := parseRow(rows)
		if err != nil {
			fmt.Println("Parse Error")
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}
func parseRow(row *sql.Rows) (User, error) {
	p := User{}
	err := row.Scan(&p.ID, &p.Type, &p.Email, &p.Phone, &p.Name, &p.CreatedAt, &p.UpdatedAt, &p.Status, &p.DeletedAt, &p.Password)
	return p, err
}
