package authenticate

import (
	"fmt"
)

// Repository User Repository
type Repository struct {
	tableName string
}

// Repo repository
var Repo = Repository{tableName: "auth_session"}

// Create new User
func (r *Repository) Create(auth *Auth) (*Auth, error) {
	str := `INSERT INTO ` + r.tableName + ` (user_id, device_id, token) values(?, ?, ?)`
	result, err := DB.Exec(str, Auth.UserID, Auth.DeviceID, Auth.Token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if id, insertErr := result.LastInsertId(); insertErr == nil {
		user.ID = int(id)
	}
	return user, nil
}
