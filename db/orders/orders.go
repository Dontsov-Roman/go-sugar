package orders

import (
	"database/sql"
	"fmt"
	"strconv"

	. "../../config"
	. "../../db"
)

// Order main struct
type Order struct {
	ID          int      `json:"ID"`
	Description string   `json:"Name"`
	Status      int      `json:"Status"`
	Prices      []int    `json:"Prices"`
	CreatedAt   NullTime `json:"CreatedAt"`
	UpdatedAt   NullTime `json:"UpdatedAt"`
	DeletedAt   NullTime `json:"DeletedAt"`
}

// Repo users repository
var Repo = Repository{tableName: Config.DB.Schema + ".orders"}

// Save entity
func (u *Order) Save() (*Order, error) {
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
func (u *Order) Validate() (bool, ValidateError) {
	return Repo.Validate(u)
}

// Delete entity
func (u *Order) Delete() bool {
	return Repo.DeleteByID(strconv.Itoa(u.ID))
}
func parseRows(rows *sql.Rows) []Order {
	var users []Order
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
func parseRow(row *sql.Rows) (Order, error) {
	p := Order{}
	err := row.Scan(&p.ID, &p.Description, &p.Status, &p.Prices, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	return p, err
}
