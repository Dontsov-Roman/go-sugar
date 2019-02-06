package main

import (
	"fmt"

	"./db/users"
)

func main() {
	user := new(users.User)
	user.Name = "Roman"
	user.Email = "dontsovroman@gmail.com"
	user.Phone = "380974885047"
	user.Status = 2
	user.Type = 3
	user.ID = 14
	// users.Create(user)
	fmt.Println(users.Update(user))
	// users.DeleteByID(8)
	fmt.Printf("%#v\n", users.GetAll())
}
