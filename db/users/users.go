package users

import (
	"database/sql"
	"fmt"
	"strconv"

	. "../../config"
	. "../../db"
	"../../db/request"
	"github.com/gin-gonic/gin"
)

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

// ValidateError struct for validate user by table
type ValidateError struct {
	ID           string
	Email        string
	Phone        string
	ErrorMessage string
}

// UserRepo User Repository
type UserRepo struct {
	tableName string
	Context   *gin.Context
}

// Repo users repository
var Repo = UserRepo{tableName: Config.DB.Schema + ".users"}

// Save entity
func (u *User) Save() (bool, error) {
	if u.ID != 0 {
		return Repo.Update(u)
	}
	_, err := Repo.Create(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Validate delegate to Repo
func (u *User) Validate() (bool, ValidateError) {
	return Repo.Validate(u)
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
	Request := request.New()
	if rows, err := Request.Select().From(r.tableName).Query(); err != nil {
		fmt.Println(err)
		return []User{}
	} else {
		return parseRows(rows)
	}
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

// Validate return bool(valid or not) and ValidateError struct
func (r *UserRepo) Validate(user *User) (bool, ValidateError) {
	valid := true
	Request := request.New()
	id := strconv.Itoa(user.ID)
	validateError := ValidateError{}
	rows, err := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "id", Operator: "=", Value: id, ConcatOperator: "OR"}).
		Where(request.Condition{Column: "email", Operator: "=", Value: user.Email, ConcatOperator: "OR"}).
		Where(request.Condition{Column: "phone", Operator: "=", Value: user.Phone, ConcatOperator: "OR"}).
		Query()
	if err == nil {
		selectedUsers := parseRows(rows)
		for i := 0; i < len(selectedUsers); i++ {
			current := selectedUsers[i]
			if current.Email == user.Email {
				validateError.Email = "User with this email already exist"
				valid = false
			}
			if current.Phone == user.Phone {
				validateError.Phone = "User with this Phone already exist"
				valid = false
			}
			if current.ID == user.ID {
				validateError.ID = "User with this ID already exist"
				valid = false
			}
		}
	} else {
		valid = false
		validateError.ErrorMessage = err.Error()
	}
	return valid, validateError
}

// Update user in DB
func (r *UserRepo) Update(user *User) (bool, error) {
	str := `UPDATE ` + r.tableName + ` SET name = ?, type = ?, status = ?, email = ?, phone = ? WHERE id = ?`
	_, err := DB.Exec(str, user.Name, user.Type, user.Status, user.Email, user.Phone, user.ID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// DeleteByID - remove user from DB
func (r *UserRepo) DeleteByID(id string) bool {
	Request := request.New()
	str, sqlErr := Request.Delete().From(r.tableName).Where(request.Condition{"id", "=", id, "OR"}).ToSQL()
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
