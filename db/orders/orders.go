package orders

import (
	"database/sql"
	"fmt"
	"strconv"

	. "go-sugar/db"
)

// Order main struct
type Order struct {
	ID          int        `json:"ID"`
	Description NullString `json:"Description"`
	Time        NullTime   `json:"Time"`
	Status      NullInt64  `json:"Status"`
	UserID      NullInt64  `json:"UserID"`
	Prices      IntArray   `json:"Prices"`
	CreatedAt   NullTime   `json:"CreatedAt"`
	UpdatedAt   NullTime   `json:"UpdatedAt"`
	DeletedAt   NullTime   `json:"DeletedAt"`
}

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
	err := row.Scan(&p.ID, &p.UserID, &p.Description, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt, &p.UserID, &p.Prices)
	return p, err
}
