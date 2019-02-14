package orders

import (
	"fmt"
	"strconv"

	. "../../db"
	"../../db/request"
	"github.com/gin-gonic/gin"
)

// Repository Orders
type Repository struct {
	tableName string
	Context   *gin.Context
}

// GetAll Orders
func (r *Repository) GetAll() []Order {
	Request := request.New(DB)
	rows, err := Request.Select().From(r.tableName).Query()
	if err != nil {
		fmt.Println(err)
		return []Order{}
	}
	return parseRows(rows)
}

// Create new Order
func (r *Repository) Create(item *Order) (*Order, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, description, status) values(?, ?, ?)`
	result, err := DB.Exec(str, item.UserID, item.Description, item.Status)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if id, insertErr := result.LastInsertId(); insertErr == nil {
		item.ID = int(id)
	}
	return item, nil
}

// Validate return bool(valid or not) and ValidateError struct
func (r *Repository) Validate(item *Order) (bool, ValidateError) {
	valid := true
	Request := request.New(DB)
	id := strconv.Itoa(item.ID)
	validateError := ValidateError{}
	rows, err := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "id", Operator: "=", Value: id, ConcatOperator: "OR"}).
		Query()
	if err == nil {
		selectedOrders := parseRows(rows)
		for i := 0; i < len(selectedOrders); i++ {
			current := selectedOrders[i]
			if current.ID == item.ID {
				validateError.ID = "Order with this ID already exist"
				validateError.AddToErrorMessage(validateError.ID)
				valid = false
			}
		}
	} else {
		valid = false
		validateError.ErrorMessage = err.Error()
	}
	return valid, validateError
}

// Update price in DB
func (r *Repository) Update(item *Order) (bool, error) {
	str := `UPDATE ` + r.tableName + ` SET user_id = ?, description = ?, status = ? WHERE id = ?`
	_, err := DB.Exec(str, item.UserID, item.Description, item.Status)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// DeleteByID - remove user from DB
func (r *Repository) DeleteByID(id string) bool {
	Request := request.New(DB)
	str, sqlErr := Request.
		Delete().
		From(r.tableName).
		Where(request.Condition{Column: "id", Operator: "=", Value: id, ConcatOperator: "OR"}).
		ToSQL()
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
