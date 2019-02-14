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
	UserID      int      `json:"UserID"`
	Prices      []int    `json:"Prices"`
	CreatedAt   NullTime `json:"CreatedAt"`
	UpdatedAt   NullTime `json:"UpdatedAt"`
	DeletedAt   NullTime `json:"DeletedAt"`
}

// Repo users repository
var Repo = Repository{tableName: Config.DB.Schema + ".orders"}

// Save entity
func (item *Order) Save() (*Order, error) {
	if item.ID != 0 {
		ok, err := Repo.Update(item)
		if !ok {
			return nil, err
		}
		return item, err
	}
	return Repo.Create(item)
}

// Validate delegate to Repo
func (item *Order) Validate() (bool, ValidateError) {
	return Repo.Validate(item)
}

// Delete entity
func (item *Order) Delete() bool {
	return Repo.DeleteByID(strconv.Itoa(item.ID))
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
