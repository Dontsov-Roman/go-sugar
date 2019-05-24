package authenticate

import (
	"database/sql"
	"errors"
	"fmt"

	. "../../db"
	"../../db/request"
)

// Repository User Repository
type Repository struct {
	tableName string
}

// Repo repository
var Repo = Repository{tableName: "auth_session"}

// Create new auth session
func (r *Repository) Create(auth *Auth) (*Auth, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, device_id, token) values(?, ?, ?)`
	_, err := DB.Exec(str, auth.UserID, auth.DeviceID, auth.Token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return auth, nil
}

// GetByDeviceID get Auth session by device_id
func (r *Repository) GetByDeviceID(DeviceID string) (*Auth, error) {
	Request := request.New(DB)
	sql, _ := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "device_id", Operator: "=", Value: DeviceID, ConcatOperator: "OR"}).
		ToSQL()
	fmt.Println(sql)
	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	auths := parseRows(rows)
	if len(auths) > 0 {
		return &auths[0], nil
	}
	return nil, errors.New("no user with this device id")
}

func parseRows(rows *sql.Rows) []Auth {
	var auths []Auth
	for rows.Next() {
		p, err := parseRow(rows)
		if err != nil {
			fmt.Println("Parse Error")
			fmt.Println(err)
			continue
		}
		auths = append(auths, p)
	}
	return auths
}
func parseRow(row *sql.Rows) (Auth, error) {
	p := Auth{}
	err := row.Scan(&p.UserID, &p.DeviceID, &p.Token, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}
