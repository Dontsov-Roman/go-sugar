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

// SaveUser with shouldBindJSON
func SaveUser(c *gin.Context) {
	user := users.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		fmt.Println(user)
		if user.Save() {
			if user.ID == 0 {
				Created(c)
			} else {
				Saved(c)
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		}
	}
}
