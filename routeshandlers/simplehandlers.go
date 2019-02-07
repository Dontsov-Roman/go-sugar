package routeshandlers

import (
	"github.com/gin-gonic/gin"
)

// GetAllNoDataJSON simple no data handler
func GetAllNoDataJSON(c *gin.Context) {
	c.JSON(404, gin.H{"msg": "No data"})
}

// Deleted Successfull
func Deleted(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "deleted"})
}
