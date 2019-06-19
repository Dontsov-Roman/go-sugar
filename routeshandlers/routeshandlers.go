package routeshandlers

import (
	"errors"
	"fmt"
	"go-sugar/db"
	"net/http"
	"strconv"
	"strings"

	"go-sugar/db/authsession"
	"go-sugar/db/orders"
	"go-sugar/db/prices"
	"go-sugar/db/request"
	"go-sugar/db/users"

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
		id := item.ID
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
		id := item.ID
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

// GetAllOrdersByUser - get all prices for main screen for authorizedUsers
func GetAllOrdersByUser(c *gin.Context) {
	var data []orders.Order
	user, err := getUser(c)
	if err != nil {
		Unauthorized(c)
		return
	}
	orderBy := c.DefaultQuery("orderBy", "time,id")
	orderType := c.DefaultQuery("orderType", "ASC")
	order := request.Order{By: strings.Split(orderBy, ","), Asc: orderType == "ASC"}
	if user.IsAdmin() {
		data = orders.Repo.GetAll(&order)
	} else {
		data = orders.Repo.GetAllByUser(&order, user)
	}
	if len(data) > 0 {
		c.JSON(200, gin.H{"data": data})
		return
	}
	GetAllNoDataJSON(c)
}

// GetAllOrders - get all prices for main screen for admin
func GetAllOrders(c *gin.Context) {
	orderBy := c.DefaultQuery("orderBy", "time,id")
	orderType := c.DefaultQuery("orderType", "ASC")
	order := request.Order{By: strings.Split(orderBy, ","), Asc: orderType == "ASC"}
	data := orders.Repo.GetAll(&order)
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
	if user, err := getUser(c); err == nil {
		if err := c.ShouldBindJSON(&item); err != nil {
			fmt.Println("shouldBindJSON", err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		} else {
			id := item.ID
			item.UserID = db.IntToNullInt(user.ID)
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
	} else {
		Unauthorized(c)
	}

}

// RegistrateByEmail handler
func RegistrateByEmail(c *gin.Context) {
	registrateByEmail := users.RegistrateByEmailUser{}
	if err := c.ShouldBindJSON(&registrateByEmail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	newUser := users.User{
		ID:       registrateByEmail.ID,
		Name:     registrateByEmail.Name,
		Email:    registrateByEmail.Email,
		Phone:    registrateByEmail.Phone,
		Password: registrateByEmail.Password,
	}
	savedItem, err := newUser.Save()
	if err == nil {
		token, err := users.Repo.CreateJWT(savedItem)
		fmt.Println(len(token))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		auth := authsession.Auth{UserID: savedItem.ID, DeviceID: registrateByEmail.DeviceID, Token: token}
		if _, authErr := auth.Save(); authErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": authErr.Error()})
			return
		}
		savedItem.Password = ""
		c.JSON(http.StatusOK, gin.H{"data": savedItem, "token": auth.Token})
		return
	}
	ok, validateError := newUser.Validate()
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": validateError.ErrorMessage, "data": validateError})
}

// GetTokenByDeviceID handler
func GetTokenByDeviceID(c *gin.Context) {
	auth, err := authsession.GetByDeviceID(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": auth.Token})
	}
}

// AuthByEmail and password
func AuthByEmail(c *gin.Context) {
	creds := users.AuthByEmail{}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	user, err := users.Repo.FindByEmail(creds.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	if user.CheckPassword(creds.Password) {
		token, err := users.Repo.CreateJWT(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		auth := authsession.Auth{UserID: user.ID, DeviceID: creds.DeviceID, Token: token}
		auth.Save()
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{"data": user, "token": token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong password"})
	}
}

// GetProfile get profile by token
func GetProfile(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		Unauthorized(c)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUser by gin.context, parsing Authorization Header;
func getUser(c *gin.Context) (*users.User, error) {
	authHeader := c.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) > 1 {
		token := splitToken[1]
		return users.Repo.ParseJWT(token)
	}
	return nil, errors.New("No token")
}

// AuthMiddleware require auth
func AuthMiddleware(c *gin.Context) {
	_, err := getUser(c)
	if err != nil {
		Unauthorized(c)
		return
	}
	c.Next()
}

// AdminMiddleware require auth and admin
func AdminMiddleware(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		Unauthorized(c)
		return
	}
	if !user.IsAdmin() {
		Forbidden(c)
		return
	}
	c.Next()
}
