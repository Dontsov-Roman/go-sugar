package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullTime new simple NullTime
type NullTime struct {
	mysql.NullTime
}

// NullString new simple NullString
type NullString struct {
	sql.NullString
}

// IntArray new struct for []int
type IntArray struct {
	String string
	Valid  bool
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// IntToNullInt helper
func IntToNullInt(i int) NullInt64 {
	var nullInt = NullInt64{}
	nullInt.Scan(i)
	return nullInt
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// ToInt - conver to int(unsafe)
func (ni *NullInt64) ToInt() int {
	return int(ni.Int64)
}

// ToString - conver to int(unsafe)
func (ni *NullInt64) ToString() string {
	return strconv.Itoa(ni.ToInt())
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC1123Z))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	if s == "null" {
		nt.Valid = false
		return nil
	}
	x, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// Scan for new struct
func (ia *IntArray) Scan(value interface{}) error {
	var i NullString
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ia = IntArray{i.String, false}
	} else {
		*ia = IntArray{i.String, true}
	}
	return nil
}

// MarshalJSON for NullTime
func (ia *IntArray) MarshalJSON() ([]byte, error) {
	if !ia.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(StringToIntArray(ia.String, ","))
}

// UnmarshalJSON for NullString
func (ia *IntArray) UnmarshalJSON(b []byte) error {
	var arr []int
	err := json.Unmarshal(b, &arr)
	if err == nil {
		ia.String = ArrayToString(arr, ",")
	}
	ia.Valid = (err == nil)
	return err
}

// StringToIntArray return []int
func StringToIntArray(str string, delimiter string) []int {
	strs := strings.Split(str, delimiter)
	arr := make([]int, len(strs))
	for i := range arr {
		arr[i], _ = strconv.Atoi(strs[i])
	}
	return arr
}

// ArrayToString return string
func ArrayToString(str []int, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(str), " ", delimiter, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(str), " "), delimiter), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(str)), delimiter), "[]")
}
