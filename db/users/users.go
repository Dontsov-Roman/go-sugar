package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-sql-driver/mysql"

	. "../../config"
	. "../../db"
	"../../db/request"
	"github.com/gin-gonic/gin"
)

type NullTime struct {
	mysql.NullTime
}

//User main struct
type User struct {
	ID        int      `json:"ID"`
	Name      string   `json:"Name"`
	Type      int      `json:"Type"`
	Status    int      `json:"Status"`
	Email     string   `json:"Email"`
	Phone     string   `json:"Phone"`
	CreatedAt NullTime `json:"CreatedAt"`
	UpdatedAt NullTime `json:"UpdatedAt"`
	DeletedAt NullTime `json:"DeletedAt"`
}

// UserRepo User Repository
type UserRepo struct {
	tableName string
	Context   *gin.Context
}

// Repo users repository
var Repo = UserRepo{tableName: Config.DB.Schema + ".users"}

// MarshalJSON for NullTime
func (ni *NullTime) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Time)
}

// Save entity
func (u *User) Save() bool {
	if u.ID != 0 {
		return Repo.Update(u)
	} else {
		_, err := Repo.Create(u)
		if err != nil {
			return false
		}
		return true
	}
}

// Delete entity
func (u *User) Delete() bool {
	return Repo.DeleteByID(strconv.Itoa(u.ID))
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
	err := row.Scan(&p.ID, &p.Type, &p.Email, &p.Phone, &p.Name, &p.CreatedAt, &p.UpdatedAt, &p.Status, &p.DeletedAt)
	return p, err
}

// GetAll Users
func (r *UserRepo) GetAll() []User {
	// str := `Select * from ` + r.tableName
	Request := request.New()
	str, sqlErr := Request.Select().From(r.tableName).ToSQL()
	if sqlErr != nil {
		return []User{}
	}
	rows, err := DB.Query(str)
	if err != nil {
		fmt.Println(err)
	}
	return parseRows(rows)
}

// Create new User
func (r *UserRepo) Create(user *User) (*User, error) {
	str := `INSERT INTO ` + r.tableName + ` (type, email, phone, name, status) values(?, ?, ?, ?, ?)`
	_, err := DB.Exec(str, user.Type, user.Email, user.Phone, user.Name, user.Status)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// user.ID = result.LastInsertId()
	return user, nil
	// var lastId int = result.LastInsertId()
	// fmt.Println(lastId)
}

// Update user in DB
func (r *UserRepo) Update(user *User) bool {
	str := `UPDATE ` + r.tableName + ` SET name = ?, type = ?, status = ?, email = ?, phone = ? WHERE id = ?`
	_, err := DB.Exec(str, user.Name, user.Type, user.Status, user.Email, user.Phone, user.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// DeleteByID - remove user from DB
func (r *UserRepo) DeleteByID(id string) bool {
	Request := request.New()
	str, sqlErr := Request.Delete().From(r.tableName).Where(request.Condition{"id", "=", id}).ToSQL()
	if sqlErr != nil {
		fmt.Println(sqlErr)
		return false
	}
	result, err := DB.Exec(str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
	return true
}
