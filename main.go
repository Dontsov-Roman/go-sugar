package main

import (
	"go-sugar/routeshandlers"

	. "go-sugar/config"

	"github.com/gin-gonic/gin"
)

// Routes main struct for routes
type Routes struct {
	Profile string
	Users   string
	Prices  string
	Orders  string
	Reserve string
}

var routes = Routes{Profile: "/profile", Users: "/users", Prices: "/prices", Orders: "/orders", Reserve: "/reserve"}

func main() {
	route := gin.Default()
	// users
	route.GET("/get-token-by-device-id/:id", routeshandlers.GetTokenByDeviceID)
	route.POST("/registrate-by-email", routeshandlers.RegistrateByEmail)
	route.POST("/auth-by-email", routeshandlers.AuthByEmail)
	authorizedProfile := route.Group(routes.Profile)
	{
		authorizedProfile.Use(routeshandlers.AuthMiddleware)
		authorizedProfile.GET("", routeshandlers.GetProfile)
		authorizedProfile.POST("", routeshandlers.SaveUser)
		authorizedProfile.PUT("", routeshandlers.SaveUser)

	}
	authorizedUsers := route.Group(routes.Users)
	{
		authorizedUsers.Use(routeshandlers.AuthMiddleware)
		authorizedUsers.GET("", routeshandlers.GetAllUsers)
		authorizedUsers.DELETE("/:id", routeshandlers.DeleteUser)
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
	route.GET(routes.Reserve, routeshandlers.GetAllReserve)
	authorizedOrders := route.Group(routes.Orders)
	{
		authorizedOrders.Use(routeshandlers.AuthMiddleware)
		authorizedOrders.GET("", routeshandlers.GetAllOrdersByUser)
		authorizedOrders.DELETE("/:id", routeshandlers.DeleteOrder)
		authorizedOrders.POST("", routeshandlers.SaveOrder)
		authorizedOrders.PUT("", routeshandlers.SaveOrder)
	}
	adminOrders := route.Group(routes.Orders)
	{
		adminOrders.Use(routeshandlers.AdminMiddleware)
		authorizedOrders.GET("/all", routeshandlers.GetAllOrders)

	}
	route.Run(":" + Config.HTTPPort)
}
