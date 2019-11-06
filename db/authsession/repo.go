package authsession

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	. "go-sugar/db"
	"go-sugar/db/request"
)

// Columns
const (
	UserID         string = "user_id"
	DeviceIDColumn string = "device_id"
	Token          string = "token"
	CreatedAt      string = "created_at"
	UpdatedAt      string = "updated_at"
)

// Repository User Repository
type Repository struct {
	tableName string
}

// Repo repository
var Repo = Repository{tableName: "auth_session"}

// CleanBeforeCreate - remove previous all user's session
func (r *Repository) CleanBeforeCreate(a *Auth) (bool, error) {
	Request := request.New(DB)
	_, err := Request.
		Delete().
		From(r.tableName).
		Where(Request.NewCond(UserID, "=", strconv.Itoa(a.UserID))).
		Where(Request.NewCond(DeviceIDColumn, "=", a.DeviceID)).
		Exec()
	if err != nil {
		fmt.Println("CleanBeforeCreate: ", err)
		return false, err
	}
	return true, nil
}

// Create new auth session
func (r *Repository) Create(auth *Auth) (*Auth, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, device_id, token) values(?, ?, ?)`
	fmt.Println(str)
	_, err := DB.Exec(str, auth.UserID, auth.DeviceID, auth.Token)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// GetByDeviceID get Auth session by device_id
func (r *Repository) GetByDeviceID(DeviceID string) (*Auth, error) {
	Request := request.New(DB)
	var orderBy []string
	orderBy = append(orderBy, CreatedAt)
	rows, err := Request.
		Select([]string{}).
		From(r.tableName).
		Where(Request.NewCond(DeviceIDColumn, "=", DeviceID)).
		OrderBy(orderBy).
		Desc().
		Limit(1).
		Query()
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
