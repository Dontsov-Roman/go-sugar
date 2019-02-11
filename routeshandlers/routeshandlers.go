package routeshandlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../db/users"
	"github.com/gin-gonic/gin"
)

// GetAllUsers - handler for route getAllUsers
func GetAllUsers(c *gin.Context) {
	data := users.Repo.GetAll()
	if len(data) > 0 {
		c.JSON(200, gin.H{"data": data})
		return
	}
	GetAllNoDataJSON(c)
}

// DeleteUser by param path
func DeleteUser(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		user := users.User{ID: id}
		if user.Delete() {
			Deleted(c)
			return
		}
	}
	GetAllNoDataJSON(c)
}

// SaveUser with ShouldBindJSON
func SaveUser(c *gin.Context) {
	user := users.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		fmt.Println(user)
		var id int = user.ID
		if savedUser, err := user.Save(); err == nil {
			if id == 0 {
				Created(c, savedUser)
			} else {
				Saved(c)
			}
		} else {
			ok, validateError := user.Validate()
			if ok {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": validateError.ErrorMessage, "data": validateError})
			}
		}
	}
}
