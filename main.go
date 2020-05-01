package main

import (
	. "go-sugar/config"
	"go-sugar/routeshandlers"

	"github.com/gin-gonic/gin"
)

// Routes main struct for routes
type Routes struct {
	Profile  string
	Users    string
	Prices   string
	Orders   string
	Reserved string
}

var routes = Routes{Profile: "/profile", Users: "/users", Prices: "/prices", Orders: "/orders", Reserved: "/reserved"}

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
	adminUsers := route.Group(routes.Users)
	{
		adminUsers.Use(routeshandlers.AdminMiddleware)
		adminUsers.POST("", routeshandlers.SaveUser)
		adminUsers.PUT("", routeshandlers.SaveUser)
	}

	// Prices
	route.GET(routes.Prices, routeshandlers.GetAllPrices)
	authorizedPrices := route.Group(routes.Prices)
	{
		authorizedPrices.Use(routeshandlers.AuthMiddleware)
		authorizedPrices.GET("/:id", routeshandlers.GetPriceByID)
		authorizedPrices.DELETE("/:id", routeshandlers.DeletePrice)
		authorizedPrices.POST("", routeshandlers.SavePrice)
		authorizedPrices.PUT("", routeshandlers.SavePrice)
	}

	// Authorized
	route.GET(routes.Reserved, routeshandlers.GetAllReserved)
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
