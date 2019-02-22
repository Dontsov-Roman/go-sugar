package ordersprices

import (
	"database/sql"
	"fmt"

	. "../../config"
	. "../../db"
)

// OrderPrice main struct
type OrderPrice struct {
	OrderID int
	UserID  NullInt64
	PriceID NullInt64
}

// Repo entity
var Repo = Repository{tableName: Config.DB.Schema + ".orders_prices"}

func parseRows(rows *sql.Rows) []OrderPrice {
	var users []OrderPrice
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
func parseRow(row *sql.Rows) (OrderPrice, error) {
	p := OrderPrice{}
	err := row.Scan(&p.OrderID, &p.UserID, &p.PriceID)
	return p, err
}
