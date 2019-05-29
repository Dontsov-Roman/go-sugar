package main

import (
	"go-sugar/routeshandlers"

	. "go-sugar/config"

	"github.com/gin-gonic/gin"
)

// Routes main struct for routes
type Routes struct {
	Users  string
	Prices string
	Orders string
}

var routes = Routes{Users: "/users", Prices: "/prices", Orders: "/orders"}

func main() {
	route := gin.Default()
	// users
	route.GET(routes.Users+"/get-token-by-device-id/:id", routeshandlers.GetTokenByDeviceID)
	route.POST(routes.Users+"/registrate-by-email", routeshandlers.RegistrateByEmail)
	route.POST(routes.Users+"/auth-by-email", routeshandlers.AuthByEmail)
	authorizedUsers := route.Group(routes.Users)
	{
		authorizedUsers.Use(routeshandlers.AuthMiddleware)
		authorizedUsers.GET("", routeshandlers.GetAllUsers)
		authorizedUsers.DELETE("/:id", routeshandlers.DeleteUser)
		authorizedUsers.POST("", routeshandlers.SaveUser)
		authorizedUsers.PUT("", routeshandlers.SaveUser)
	}

	// Prices
	route.GET(routes.Prices, routeshandlers.GetAllPrices)
	authorizedPrices := route.Group(routes.Prices)
	{
		authorizedPrices.Use(routeshandlers.AuthMiddleware)
		authorizedPrices.DELETE("/:id", routeshandlers.DeletePrice)
		authorizedPrices.POST("", routeshandlers.SavePrice)
		authorizedPrices.PUT("", routeshandlers.SavePrice)
	}

	// Authorized
	authorizedOrders := route.Group(routes.Orders)
	{
		authorizedOrders.Use(routeshandlers.AuthMiddleware)
		authorizedOrders.GET("", routeshandlers.GetAllOrders)
		authorizedOrders.DELETE("/:id", routeshandlers.DeleteOrder)
		authorizedOrders.POST("", routeshandlers.SaveOrder)
		authorizedOrders.PUT("", routeshandlers.SaveOrder)
	}

	route.Run(":" + Config.HTTPPort)
}
