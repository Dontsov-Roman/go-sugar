package ordersprices

import (
	"fmt"
	"strconv"

	. "../../db"
	"../../db/request"
	"github.com/gin-gonic/gin"
)

// Repository OrderPrices
type Repository struct {
	tableName string
	Context   *gin.Context
}

// GetAll OrderPrices
func (r *Repository) GetAll() []OrderPrice {
	Request := request.New(DB)
	rows, err := Request.Select().From(r.tableName).Query()
	if err != nil {
		fmt.Println(err)
		return []OrderPrice{}
	}
	return parseRows(rows)
}

// Create new OrderPrice
func (r *Repository) Create(items []OrderPrice) ([]OrderPrice, error) {
	Request := request.New(DB)
	keys := []string{"order_id", "user_id", "price_id"}
	var values = [][]string{}
	for _, item := range items {
		values = append(values, []string{strconv.Itoa(item.OrderID), strconv.Itoa(item.UserID), strconv.Itoa(item.PriceID)})
	}
	_, err := Request.Insert().
		Into(r.tableName).
		Values(keys, values).
		Exec()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

// Validate return bool(valid or not) and ValidateError struct
func (r *Repository) Validate(item *OrderPrice) (bool, ValidateError) {
	valid := true
	Request := request.New(DB)
	id := strconv.Itoa(item.OrderID)
	priceID := strconv.Itoa(item.PriceID)
	validateError := ValidateError{}
	rows, err := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "order_id", Operator: "=", Value: id, ConcatOperator: "AND"}).
		Where(request.Condition{Column: "price_id", Operator: "=", Value: priceID, ConcatOperator: "AND"}).
		Query()
	if err == nil {
		selectedOrderPrices := parseRows(rows)
		if len(selectedOrderPrices) > 0 {
			validateError.OrderIDPriceID = "OrderPrice with this OrderID and PriceID already exist"
			validateError.AddToErrorMessage(validateError.OrderIDPriceID)
			valid = false
		}
	} else {
		valid = false
		validateError.ErrorMessage = err.Error()
	}
	return valid, validateError
}

// DeleteByOrderID - remove user from DB
func (r *Repository) DeleteByOrderID(id string) bool {
	Request := request.New(DB)
	str, sqlErr := Request.
		Delete().
		From(r.tableName).
		Where(request.Condition{Column: "order_id", Operator: "=", Value: id, ConcatOperator: "OR"}).
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

// GetByOrderID - get all by IrderID from DB
func (r *Repository) GetByOrderID(id string) []OrderPrice {
	Request := request.New(DB)
	rows, err := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "order_id", Operator: "=", Value: id, ConcatOperator: "OR"}).
		Query()
	if err != nil {
		fmt.Println(err)
		return []OrderPrice{}
	}
	return parseRows(rows)
}
