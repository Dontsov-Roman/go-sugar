package orders

import (
	"fmt"
	"strconv"

	. "go-sugar/config"
	. "go-sugar/db"
	"go-sugar/db/ordersprices"
	"go-sugar/db/request"
	"go-sugar/db/users"

	"github.com/gin-gonic/gin"
)

// Columns
const (
	ID          string = "id"
	UserID      string = "user_id"
	Description string = "description"
	Time        string = "time"
	TimeEnd     string = "time_end"
	Status      string = "status"
	CreatedAt   string = "created_at"
	UpdatedAt   string = "updated_at"
	DeletedAt   string = "deleted_at"
)

// Repository Orders
type Repository struct {
	tableName string
	Context   *gin.Context
	delimiter string
}

// Repo users repository
var Repo = Repository{tableName: Config.DB.Schema + ".orders", delimiter: ","}

// GetAllByUser Orders by User
func (r *Repository) GetAllByUser(o *request.Order, u *users.User) []Order {
	o.TablePrefix = r.tableName
	userID := strconv.Itoa(u.ID)
	str := "SELECT *, (select group_concat(price_id) FROM orders_prices WHERE order_id=orders.id) AS prices FROM " + r.tableName + " WHERE user_id = " + userID + " ORDER BY " + o.ToString(true)
	rows, err := DB.Query(str)
	if err != nil {
		return []Order{}
	}
	orders := parseRows(rows)
	return orders
}

// GetAll Orders User
func (r *Repository) GetAll(o *request.Order) []Order {
	o.TablePrefix = r.tableName
	str := "SELECT *, (select group_concat(price_id) FROM orders_prices WHERE order_id=orders.id) AS prices FROM " + r.tableName + " ORDER BY " + o.ToString(true)
	rows, err := DB.Query(str)
	if err != nil {
		return []Order{}
	}
	orders := parseRows(rows)
	return orders
}

// GetAllReserved return an array of Reserve
func (r *Repository) GetAllReserved() []Reserve {
	Request := request.New(DB)
	var columns []string
	columns = append(columns, ID, Time, TimeEnd)
	condition := Request.NewCondition(TimeEnd, ">", "NOW()", "OR", true)
	req := Request.Select(columns).From(r.tableName).Where(condition)
	rows, err := req.Query()
	if err != nil {
		fmt.Println(err)
		return []Reserve{}
	}
	return parseRowsReserve(rows)
}

// Create new Order
func (r *Repository) Create(item *Order) (*Order, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, description, status, time, time_end) values(?, ?, ?, ?, ?)`
	result, err := DB.Exec(str, item.UserID, item.Description, item.Status, item.Time, item.TimeEnd)
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
		Select([]string{}).
		From(r.tableName).
		Where(Request.NewCond(ID, "=", id)).
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
		Where(Request.NewCond(ID, "=", id)).
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
