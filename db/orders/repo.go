package orders

import (
	"fmt"
	"strconv"

	. "go-sugar/config"
	. "go-sugar/db"
	"go-sugar/db/ordersprices"
	"go-sugar/db/request"

	"github.com/gin-gonic/gin"
)

// Repository Orders
type Repository struct {
	tableName string
	Context   *gin.Context
	delimiter string
}

// Repo users repository
var Repo = Repository{tableName: Config.DB.Schema + ".orders", delimiter: ","}

// GetAll Orders
func (r *Repository) GetAll(o *request.Order) []Order {
	o.TablePrefix = r.tableName
	str := "SELECT *,(SELECT user_id FROM orders_prices where order_id=orders.id limit 1) as user_id, (select group_concat(price_id) FROM orders_prices WHERE order_id=orders.id) AS prices FROM " + r.tableName + " ORDER BY " + o.ToString(true)
	fmt.Println(str)
	rows, err := DB.Query(str)
	if err != nil {
		fmt.Println(err)
		return []Order{}
	}
	orders := parseRows(rows)
	return orders
}

// Create new Order
func (r *Repository) Create(item *Order) (*Order, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, description, status, time) values(?, ?, ?, ?)`
	result, err := DB.Exec(str, item.UserID, item.Description, item.Status, item.Time)
	var OP []ordersprices.OrderPrice
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if id, insertErr := result.LastInsertId(); insertErr == nil {
		item.ID = int(id)
	}
	for _, priceID := range StringToIntArray(item.Prices.String, r.delimiter) {
		OP = append(OP, ordersprices.OrderPrice{OrderID: item.ID, UserID: item.UserID, PriceID: IntToNullInt(priceID)})
	}
	ordersprices.Repo.Create(OP)
	return item, nil
}

// Update price in DB
func (r *Repository) Update(item *Order) (bool, error) {
	str := `UPDATE ` + r.tableName + ` SET description = ?, status = ? WHERE id = ?`
	_, err := DB.Exec(str, item.Description, item.Status)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
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
	_, err := DB.Exec(str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	// fmt.Println(result.RowsAffected()) // количество затронутых строк
	return true
}
