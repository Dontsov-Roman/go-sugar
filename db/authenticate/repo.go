package authenticate

import "fmt"

// Repository User Repository
type Repository struct {
	tableName string
}

// Repo repository
var Repo = Repository{tableName: "authenticate"}

// Create new User
func (r *Repository) Create(auth *Auth) (*Auth, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, email, phone, name, status, password) values(?, ?, ?, ?, ?, ?)`
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
