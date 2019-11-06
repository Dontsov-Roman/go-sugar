package prices

import (
	"fmt"
	"strconv"
	"strings"

	. "go-sugar/db"
	"go-sugar/db/request"

	"github.com/gin-gonic/gin"
)

// Columns
const (
	ID          string = "id"
	Name        string = "name"
	Status      string = "status"
	PriceColumn string = "price"
	Time        string = "time"
	CreatedAt   string = "created_at"
	UpdatedAt   string = "updated_at"
	DeletedAt   string = "deleted_at"
)

// Repository Price
type Repository struct {
	tableName string
	Context   *gin.Context
}

// GetAll Prices
func (r *Repository) GetAll() []Price {
	Request := request.New(DB)
	rows, err := Request.Select([]string{}).From(r.tableName).Query()
	if err != nil {
		fmt.Println(err)
		return []Price{}
	}
	return parseRows(rows)
}

// Create new Price
func (r *Repository) Create(item *Price) (*Price, error) {
	columns := []string{Name, Status, PriceColumn, Time}
	str := `INSERT INTO ` + r.tableName + ` (` + strings.Join(columns, ",") + `) values(?, ?, ?, ?)`
	result, err := DB.Exec(str, item.Name, item.Status, item.Price, item.Time)
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
func (r *Repository) Validate(item *Price) (bool, ValidateError) {
	valid := true
	Request := request.New(DB)
	id := strconv.Itoa(item.ID)
	validateError := ValidateError{}
	rows, err := Request.
		Select([]string{}).
		From(r.tableName).
		Where(Request.NewCond(ID, "=", id)).
		Query()
	if err == nil {
		selectedPrices := parseRows(rows)
		for i := 0; i < len(selectedPrices); i++ {
			current := selectedPrices[i]
			if current.ID == item.ID {
				validateError.ID = "Price with this ID already exist"
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
func (r *Repository) Update(item *Price) (bool, error) {
	str := `UPDATE ` + r.tableName + ` SET name = ?, status = ?, price = ?, time = ? WHERE id = ?`
	_, err := DB.Exec(str, item.Name, item.Status, item.Price, item.Time, item.ID)
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
		Where(Request.NewCond(ID, "=", id)).
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
