package routeshandlers

import (
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
	id := c.Param("id")
	if users.Repo.DeleteByID(id) {
		Deleted(c)
		return
	}
	GetAllNoDataJSON(c)
}
