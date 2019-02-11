package routeshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllNoDataJSON simple no data handler
func GetAllNoDataJSON(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"msg": "No data"})
}

// Saved Simple saved status
func Saved(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Saved"})
}

// Created simple created status
func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}

// Deleted Successfull
func Deleted(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "deleted"})
}
