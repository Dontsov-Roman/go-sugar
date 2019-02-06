package users

import (
	"database/sql"
	"fmt"

	. "../../config"
	. "../../db"
)

var dbName string = Config.DB.Schema + ".users"

//User main struct
type User struct {
	ID        int
	Name      string
	Type      int
	Status    int
	Email     string
	Phone     string
	CreatedAt string
	UpdatedAt string
}

func parseRows(rows *sql.Rows) []User {
	var users []User
	for rows.Next() {
		p, err := parseRow(rows)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}
func parseRow(row *sql.Rows) (User, error) {
	p := User{}
	err := row.Scan(&p.ID, &p.Type, &p.Email, &p.Phone, &p.Name, &p.CreatedAt, &p.UpdatedAt, &p.Status)
	return p, err
}

// GetAll Users
func GetAll() []User {
	str := `Select * from ` + dbName
	rows, err := DB.Query(str)
	if err != nil {
		fmt.Println(err)
	}
	return parseRows(rows)
}

// Create new User
func Create(user *User) *User {
	str := `INSERT INTO ` + dbName + ` (type, email, phone, name, status) values(?, ?, ?, ?, ?)`
	result, err := DB.Exec(str, user.Type, user.Email, user.Phone, user.Name, user.Status)
	if err != nil {
		fmt.Println(err)
	}
	// user.ID = result.LastInsertId()
	fmt.Println(result.LastInsertId())
	return user
	// var lastId int = result.LastInsertId()
	// fmt.Println(lastId)
}

// Update user in DB
func Update(user *User) bool {
	str := `UPDATE ` + dbName + ` SET name = ?, type = ?, status = ?, email = ?, phone = ? WHERE id = ?`
	_, err := DB.Exec(str, user.Name, user.Type, user.Status, user.Email, user.Phone, user.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// DeleteByID - remove user from DB
func DeleteByID(id int) {
	str := `DELETE FROM ` + dbName + ` WHERE id = ?`
	result, err := DB.Exec(str, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
