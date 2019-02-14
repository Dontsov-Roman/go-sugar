package prices

import (
	"database/sql"
	"fmt"
	"strconv"

	. "../../config"
	. "../../db"
)

//Price main struct
type Price struct {
	ID        int      `json:"ID"`
	Name      string   `json:"Name"`
	Status    int      `json:"Status"`
	Price     int      `json:"Price"`
	Time      int      `json:Time` // in minutes
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
	DeletedAt NullTime `json:"DeletedAt"`
}

// Repo users repository
var Repo = PriceRepo{tableName: Config.DB.Schema + ".prices"}

// Save entity
func (u *Price) Save() (*Price, error) {
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
func (u *Price) Validate() (bool, ValidateError) {
	return Repo.Validate(u)
}

// Delete entity
func (u *Price) Delete() bool {
	return Repo.DeleteByID(strconv.Itoa(u.ID))
}
func parseRows(rows *sql.Rows) []Price {
	var users []Price
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
func parseRow(row *sql.Rows) (Price, error) {
	p := Price{}
	err := row.Scan(&p.ID, &p.Name, &p.Status, &p.Price, &p.Time, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	return p, err
}
