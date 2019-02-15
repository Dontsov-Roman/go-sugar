package users

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	. "../../config"
	. "../../db"
	"../../db/request"
)

// Repository User Repository
type Repository struct {
	tableName      string
	salt           string
	secretForToken string
	Context        *gin.Context
}

// Repo users repository
var Repo = Repository{tableName: Config.DB.Schema + ".users", salt: "sweet_sugar_67n334g6", secretForToken: "sweet_sugar_346436bb43gh463"}

// GetAll Users
func (r *Repository) GetAll() []User {
	Request := request.New(DB)
	rows, err := Request.Select().From(r.tableName).Query()
	if err != nil {
		fmt.Println(err)
		return []User{}
	}
	return parseRows(rows)
}

// Create new User
func (r *Repository) Create(user *User) (*User, error) {
	user.Password = r.CreateHash(user.Password)
	str := `INSERT INTO ` + r.tableName + ` (type, email, phone, name, status, password) values(?, ?, ?, ?, ?, ?)`
	result, err := DB.Exec(str, user.Type, user.Email, user.Phone, user.Name, user.Status, user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if id, insertErr := result.LastInsertId(); insertErr == nil {
		user.ID = int(id)
	}
	return user, nil
}

// Validate return bool(valid or not) and ValidateError struct
func (r *Repository) Validate(user *User) (bool, ValidateError) {
	valid := true
	Request := request.New(DB)
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
			if current.ID == user.ID {
				validateError.ID = "User with this ID already exist"
				validateError.AddToErrorMessage(validateError.ID)
				valid = false
			}
			if current.Email == user.Email {
				validateError.Email = "User with this email already exist"
				validateError.AddToErrorMessage(validateError.Email)
				valid = false
			}
			if current.Phone == user.Phone {
				validateError.Phone = "User with this Phone already exist"
				validateError.AddToErrorMessage(validateError.Phone)
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
func (r *Repository) Update(user *User) (bool, error) {
	str := `UPDATE ` + r.tableName + ` SET name = ?, type = ?, status = ?, email = ?, phone = ? WHERE id = ?`
	_, err := DB.Exec(str, user.Name, user.Type, user.Status, user.Email, user.Phone, user.ID)
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
	fmt.Println(str)
	result, err := DB.Exec(str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
	return true
}

// CreateHash return a hashed string
func (r *Repository) CreateHash(str string) string {
	aStringToHash := []byte(str + r.salt)
	sha1Bytes := sha1.Sum(aStringToHash)
	encodedStr := hex.EncodeToString(sha1Bytes[:])
	fmt.Println(str + ": " + encodedStr)
	return encodedStr
}

// ParseToken return a User
func (r *Repository) ParseToken(tokenString string) (User, error) {
	user := User{}
	mapToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return r.secretForToken, nil
	})
	if claims, ok := mapToken.Claims.(jwt.MapClaims); ok && mapToken.Valid {
		fmt.Println(claims)
		// user.ID = claims["ID"]
		// user.Name = claims["Name"]
		// user.Email = claims["Email"]
		// user.Phone = claims["Phone"]
		// user.Status = claims["Status"]
		// user.Type = claims["Type"]
	} else {
		fmt.Println(err)
	}

	return user, err
}

// CreateToken return a JWT
func (r *Repository) CreateToken(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":     u.ID,
		"Name":   u.Name,
		"Email":  u.Email,
		"Phone":  u.Phone,
		"Status": u.Status,
		"Type":   u.Type,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(r.secretForToken)
	return tokenString, err
}
