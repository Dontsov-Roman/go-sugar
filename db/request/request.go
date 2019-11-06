package request

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// requestType - enum of types
type requestType int

// Requests
const (
	UNDEFINED requestType = 0
	SELECT    requestType = 1
	UPDATE    requestType = 2
	INSERT    requestType = 3
	DELETE    requestType = 4
)

// Request structure
type Request struct {
	tableName   string
	join        string
	columns     []string
	keys        []string
	values      [][]string
	set         map[string]string
	where       []*Condition
	requestType requestType // SELECT UPDATE INSERT DELETE
	orderBy     []string
	orderAsc    bool
	offset      int
	limit       int
	db          *sql.DB
}

// Order uses for request
type Order struct {
	TablePrefix string
	By          []string
	Asc         bool
}

// Paginator uses for request limit, offset
type Paginator struct {
	Page    int
	PerPage int
}

// IRequestBuilder main interface
type IRequestBuilder interface {
	Select() *Request
	Update() *Request
	Insert() *Request
	Delete() *Request
	SetType(int) *Request
	From(string) *Request
	Join(string) *Request
	Set() *Request
	Values() *Request
	Where(string) *Request
	OrderBy([]string) *Request
	Asc() *Request
	Desc() *Request
	Offset(int) *Request
	Limit(int) *Request
	ToSQL() string
}

// ToString Order struct
func (o *Order) ToString(withPrefix bool) string {
	var str string
	var orderType string
	if o.Asc {
		orderType = "ASC"
	} else {
		orderType = "DESC"
	}
	if withPrefix {
		var stringsWithPrefix []string
		for _, v := range o.By {
			stringsWithPrefix = append(stringsWithPrefix, o.TablePrefix+"."+v)
		}
		str = strings.Join(stringsWithPrefix, ", ") + " " + orderType
	} else {
		str = strings.Join(o.By, ", ") + " " + orderType
	}
	return str
}

// New Request
func New(db *sql.DB) Request {
	request := Request{}
	request.SetDB(db)
	return request
}

// SetDB to request
func (r *Request) SetDB(db *sql.DB) *Request {
	r.db = db
	return r
}

// Select === SetType(SELECT)
func (r *Request) Select(columns []string) *Request {
	r.SetType(SELECT)
	r.columns = columns
	return r
}

// Update === SetType(UPDATE)
func (r *Request) Update(tableName string) *Request {
	r.SetType(UPDATE)
	r.tableName = tableName
	return r
}

// Insert === SetType(INSERT)
func (r *Request) Insert() *Request {
	r.SetType(INSERT)
	return r
}

// Delete === SetType(DELETE)
func (r *Request) Delete() *Request {
	r.SetType(DELETE)
	return r
}

// SetType SELECT INSERT UPDATE DELETE
func (r *Request) SetType(typeRequest requestType) *Request {
	r.requestType = typeRequest
	return r
}

// Into tableName
func (r *Request) Into(tableName string) *Request {
	r.tableName = tableName
	return r
}

// From set tableName
func (r *Request) From(tableName string) *Request {
	r.tableName = tableName
	return r
}

// Join add join string
func (r *Request) Join(join string) *Request {
	r.join = join
	return r
}

// Set the same as Values
func (r *Request) Set(key string, val string) *Request {
	r.set[key] = val
	return r
}

// Values add value to map[string]string
func (r *Request) Values(keys []string, values [][]string) *Request {
	r.keys = keys
	r.values = values
	return r
}

// Where add condition to array
func (r *Request) Where(cond *Condition) *Request {
	if cond.Column == "" || cond.Value == "" {
		return r
	}
	if cond.Operator == "" {
		cond.Operator = "="
	}
	if cond.ConcatOperator == "" {
		cond.ConcatOperator = "AND"
	}
	r.where = append(r.where, cond)
	return r
}

func (r *Request) parseWhere() (string, error) {
	str := " WHERE "
	length := len(r.where)
	fmt.Println(r.where, length)
	if length > 0 {
		for i := 0; i < length; i++ {
			str = str + r.where[i].Column + r.where[i].Operator
			if r.where[i].Native {
				str = str + r.where[i].Value
			} else {
				str = str + "\"" + r.where[i].Value + "\""
			}
			if i+1 < length {
				str = str + " " + r.where[i].ConcatOperator + " "
			}
		}
		return str, nil
	}
	return "", errors.New("No conditions in where")
}

// Order use Order struct,
func (r *Request) Order(o *Order) *Request {
	r.orderBy = o.By
	r.orderAsc = o.Asc
	return r
}

// OrderBy set order for select
func (r *Request) OrderBy(str []string) *Request {
	r.orderBy = str
	r.orderAsc = true
	return r
}

// Asc set order for select
func (r *Request) Asc() *Request {
	r.orderAsc = true
	return r
}

// Desc set order for select
func (r *Request) Desc() *Request {
	r.orderAsc = false
	return r
}

// Offset set Offset for select
func (r *Request) Offset(offset int) *Request {
	r.offset = offset
	return r
}

// Limit set limit for select
func (r *Request) Limit(limit int) *Request {
	r.limit = limit
	return r
}
func (r *Request) getColumns() string {
	if len(r.columns) < 1 {
		return "*"
	} else {
		return strings.Join(r.columns, ",")
	}
}

// ToSQL return a SQL string
func (r *Request) ToSQL() (string, error) {
	var str string
	if r.tableName == "" {
		return "", errors.New("no table name")
	}
	if r.requestType == UNDEFINED {
		return "", errors.New("no requestType")
	}
	switch r.requestType {
	case SELECT:
		str = "SELECT " + r.getColumns() + " FROM " + r.tableName
		if where, err := r.parseWhere(); err == nil {
			str = str + where
		}
		if len(r.join) > 0 {
			str += r.join
		}
		if len(r.orderBy) > 0 {
			str = str + " ORDER BY " + strings.Join(r.orderBy, ",")
			orderType := " ASC"
			if !r.orderAsc {
				orderType = " DESC"
			}
			str = str + orderType
		}

		if r.limit != 0 {
			str = str + " LIMIT " + strconv.Itoa(r.limit)
			if r.offset != 0 {
				str = str + " OFFSET " + strconv.Itoa(r.offset)
			}
		}
	case UPDATE:
		str = "UPDATE " + r.tableName + " SET "
		for key, val := range r.set {
			str = str + key + " = " + val + ", "
		}
		str = rCut(str, 2)
		if where, err := r.parseWhere(); err == nil {
			str = str + where
		}
	case INSERT:
		str = "INSERT INTO " + r.tableName
		var values []string
		for _, val := range r.values {
			values = append(values, "("+strings.Join(val, ",")+")")
		}
		str = str + " (" + strings.Join(r.keys, ",") + ") VALUES " + strings.Join(values, ",")
	case DELETE:
		str = "DELETE FROM " + r.tableName
		where, err := r.parseWhere()
		if err == nil {
			str = str + where
		} else {
			return "", err
		}
	}

	return str, nil
}

// Exec delegate to sql
func (r *Request) Exec() (sql.Result, error) {
	str, err := r.ToSQL()
	if err == nil {
		return r.db.Exec(str)
	}
	return nil, err
}

// Query delegate to sql
func (r *Request) Query() (*sql.Rows, error) {
	str, err := r.ToSQL()
	if err == nil {
		return r.db.Query(str)
	}
	return nil, err
}

func cut(text string, limit int) string {
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit])
	}
	return text
}

func rCut(text string, limit int) string {
	return cut(text, len(text)-limit)
}
