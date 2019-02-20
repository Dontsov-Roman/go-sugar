package main

import (
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
	// user := users.User{ID: 50, Name: "Roman"}
	// token, err := users.Repo.CreateJWT(&user)
	// parsedUser, err := users.Repo.ParseJWT(token)
	// fmt.Println(parsedUser, err)

	route := gin.Default()
	route.GET(routes.Users, routeshandlers.GetAllUsers)
	route.DELETE(routes.Users+"/:id", routeshandlers.DeleteUser)
	route.POST(routes.Users, routeshandlers.SaveUser)
	route.PUT(routes.Users, routeshandlers.SaveUser)

	route.GET(routes.Prices, routeshandlers.GetAllPrices)
	route.DELETE(routes.Prices+"/:id", routeshandlers.DeletePrice)
	route.POST(routes.Prices, routeshandlers.SavePrice)
	route.PUT(routes.Prices, routeshandlers.SavePrice)

	authorized := route.Group("/")
	authorized.Use(routeshandlers.AuthMiddleware)
	{
		route.GET(routes.Orders, routeshandlers.GetAllOrders)
		route.DELETE(routes.Orders+"/:id", routeshandlers.DeleteOrder)
		route.POST(routes.Orders, routeshandlers.SaveOrder)
		route.PUT(routes.Orders, routeshandlers.SaveOrder)
	}

	route.Run(":9000")
}
