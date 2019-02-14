package request

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// Request structure
type Request struct {
	tableName   string
	join        string
	values      map[string]string
	where       []Condition
	requestType int // select - 1 | update - 2 | insert - 3 | delete - 4
	orderBy     []string
	orderAsc    bool
	offset      int
	limit       int
	db          *sql.DB
}

// Condition for Where method
type Condition struct {
	Column         string
	Operator       string
	Value          string
	ConcatOperator string
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

// Select === SetType(1)
func (r *Request) Select() *Request {
	r.SetType(1)
	return r
}

// Update === SetType(2)
func (r *Request) Update() *Request {
	r.SetType(2)
	return r
}

// Insert === SetType(3)
func (r *Request) Insert() *Request {
	r.SetType(3)
	return r
}

// Delete === SetType(4)
func (r *Request) Delete() *Request {
	r.SetType(4)
	return r
}

// SetType select - 1 | update - 2 | insert - 3 | delete - 4
func (r *Request) SetType(typeRequest int) *Request {
	r.requestType = typeRequest
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
	return r.Values(key, val)
}

// Values add value to map[string]string
func (r *Request) Values(key string, val string) *Request {
	r.values[key] = val
	return r
}

// Where add condition to array
func (r *Request) Where(cond Condition) *Request {
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
			fmt.Println(length, i)
			str = str + r.where[i].Column + r.where[i].Operator + r.where[i].Value
			if i+1 < length {
				str = str + " " + r.where[i].ConcatOperator + " "
			}
		}
		return str, nil
	}
	return "", errors.New("No conditions in where")
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

// ToSQL return a SQL string
func (r *Request) ToSQL() (string, error) {
	var str string
	if r.tableName == "" {
		return "", errors.New("no table name")
	}
	if r.requestType == 0 {
		return "", errors.New("no requestType")
	}
	switch r.requestType {
	case 1:
		str = "SELECT * FROM " + r.tableName
		if where, err := r.parseWhere(); err == nil {
			str = str + where
		}

		if r.limit != 0 {
			str = str + " LIMIT " + string(r.limit)
			if r.offset != 0 {
				str = str + " OFFSET " + string(r.limit)
			}
		}
		if len(r.orderBy) > 0 {
			str = str + " ORDER BY " + strings.Join(r.orderBy, ",")
		}
	case 2:
		str = "UPDATE " + r.tableName + " SET "
		for key, val := range r.values {
			str = str + key + " = " + val + ", "
		}
		str = rCut(str, 2)
		if where, err := r.parseWhere(); err == nil {
			str = str + where
		}
	case 3:
		str = "INSERT INTO " + r.tableName
		var keys string
		var values string
		for key, val := range r.values {
			keys = keys + key + ", "
			values = values + val + ", "
		}
		keys = rCut(keys, 2)
		values = rCut(values, 2)
		str = str + " (" + keys + ") VALUES (" + values + ")"
	case 4:
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
