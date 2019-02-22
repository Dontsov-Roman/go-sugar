package main

import (
	. "./config"
	"./routeshandlers"
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
	// user := users.Repo.FindByID("50")
	// if user != nil {
	// 	token, err := users.Repo.CreateJWT(user)
	// 	parsedUser, err := users.Repo.ParseJWT(token)
	// 	fmt.Println(token, parsedUser, err)
	// }

	route := gin.Default()
	route.GET(routes.Users, routeshandlers.GetAllUsers)
	route.DELETE(routes.Users+"/:id", routeshandlers.DeleteUser)
	route.POST(routes.Users, routeshandlers.SaveUser)
	route.PUT(routes.Users, routeshandlers.SaveUser)

	route.GET(routes.Prices, routeshandlers.GetAllPrices)
	route.DELETE(routes.Prices+"/:id", routeshandlers.DeletePrice)
	route.POST(routes.Prices, routeshandlers.SavePrice)
	route.PUT(routes.Prices, routeshandlers.SavePrice)
	route.GET(routes.Orders, routeshandlers.GetAllOrders)

	authorized := route.Group(routes.Orders)
	{
		authorized.Use(routeshandlers.AuthMiddleware)

		authorized.DELETE("/:id", routeshandlers.DeleteOrder)
		authorized.POST("", routeshandlers.SaveOrder)
		authorized.PUT("", routeshandlers.SaveOrder)
	}

	route.Run(":" + Config.HTTPPort)
}
