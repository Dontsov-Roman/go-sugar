package routeshandlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../db/orders"
	"../db/prices"
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
	item := users.User{}

	if err := c.ShouldBindJSON(&item); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		var id int = item.ID
		if savedItem, err := item.Save(); err == nil {
			if id == 0 {
				Created(c, savedItem)
			} else {
				Saved(c)
			}
		} else {
			ok, validateError := item.Validate()
			if ok {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": validateError.ErrorMessage, "data": validateError})
			}
		}
	}
}

// GetAllPrices - get all prices for main screen
func GetAllPrices(c *gin.Context) {
	data := prices.Repo.GetAll()
	if len(data) > 0 {
		c.JSON(200, gin.H{"data": data})
		return
	}
	GetAllNoDataJSON(c)
}

// DeletePrice by param path
func DeletePrice(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		item := prices.Price{ID: id}
		if item.Delete() {
			Deleted(c)
			return
		}
	}
	GetAllNoDataJSON(c)
}

// SavePrice with ShouldBindJSON
func SavePrice(c *gin.Context) {
	item := prices.Price{}

	if err := c.ShouldBindJSON(&item); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		var id int = item.ID
		if savedItem, err := item.Save(); err == nil {
			if id == 0 {
				Created(c, savedItem)
			} else {
				Saved(c)
			}
		} else {
			ok, validateError := item.Validate()
			if ok {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": validateError.ErrorMessage, "data": validateError})
			}
		}
	}
}

// GetAllOrders - get all prices for main screen
func GetAllOrders(c *gin.Context) {
	data := orders.Repo.GetAll()
	if len(data) > 0 {
		c.JSON(200, gin.H{"data": data})
		return
	}
	GetAllNoDataJSON(c)
}

// DeleteOrder by param path
func DeleteOrder(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		item := orders.Order{ID: id}
		if item.Delete() {
			Deleted(c)
			return
		}
	}
	GetAllNoDataJSON(c)
}

// SaveOrder with ShouldBindJSON
func SaveOrder(c *gin.Context) {
	item := orders.Order{}

	if err := c.ShouldBindJSON(&item); err != nil {
		fmt.Println("shouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		var id int = item.ID
		if savedItem, err := item.Save(); err == nil {
			if id == 0 {
				Created(c, savedItem)
			} else {
				Saved(c)
			}
		} else {
			ok, validateError := item.Validate()
			if ok {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": validateError.ErrorMessage, "data": validateError})
			}
		}
	}
}

// AuthMiddleware require auth
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) > 1 {
		token := splitToken[1]
		_, err := users.Repo.ParseJWT(token)
		if err != nil {
			Unauthorized(c)
			return
		}
	} else {
		Unauthorized(c)
		return
	}
	fmt.Println("Passed")
	c.Next()
}
