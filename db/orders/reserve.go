package orders

import (
	"database/sql"
	"fmt"
	. "go-sugar/db"
)

// Reserve - main struct
type Reserve struct {
	ID      int      `json:"ID"`
	Time    NullTime `json:"Time"`
	TimeEnd NullTime `json:"TimeEnd"`
}

func parseRowsReserve(rows *sql.Rows) []Reserve {
	var data []Reserve
	for rows.Next() {
		p, err := parseRowReserve(rows)
		if err != nil {
			fmt.Println("Parse Error")
			fmt.Println(err)
			continue
		}
		data = append(data, p)
	}
	return data
}
func parseRowReserve(row *sql.Rows) (Reserve, error) {
	p := Reserve{}
	err := row.Scan(&p.ID, &p.Time, &p.TimeEnd)
	return p, err
}
