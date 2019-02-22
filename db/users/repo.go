package users

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"

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
	expiresAt      int64
}

// Repo users repository
var Repo = Repository{
	tableName:      Config.DB.Schema + ".users",
	salt:           "sweet_sugar_67n334g6",
	secretForToken: "sweet_sugar_346436bb43gh463",
	expiresAt:      60 * 60,
}

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
	_, err := DB.Exec(str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// FindByID - find user by ID
func (r *Repository) FindByID(id string) *User {
	Request := request.New(DB)
	rows, err := Request.
		Select().
		From(r.tableName).
		Where(request.Condition{Column: "id", Operator: "=", Value: id, ConcatOperator: "OR"}).
		Query()
	users := parseRows(rows)
	if err == nil && len(users) > 0 {
		return &users[0]
	}
	return nil
}

// CreateHash return a hashed string
func (r *Repository) CreateHash(str string) string {
	aStringToHash := []byte(str + r.salt)
	sha1Bytes := sha1.Sum(aStringToHash)
	encodedStr := hex.EncodeToString(sha1Bytes[:])
	return encodedStr
}

// GetClaims return new claims with user
func (r *Repository) GetClaims(user *User) CustomClaims {
	claims := CustomClaims{User: user}
	claims.ExpiresAt = jwt.TimeFunc().Unix() + Repo.expiresAt
	return claims
}

// CreateJWT return a JWT
func (r *Repository) CreateJWT(u *User) (string, error) {
	claims := r.GetClaims(u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(r.secretForToken))
	return tokenString, err
}

// ParseJWT return a User
func (r *Repository) ParseJWT(tokenString string) (*User, error) {
	claims := CustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(r.secretForToken), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.User, nil
}
