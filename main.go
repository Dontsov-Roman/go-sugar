package main

import (
	"./routeshandlers"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", routeshandlers.GetAllUsers)
	route.DELETE("/users/:id", routeshandlers.DeleteUser)
	route.POST("/users", routeshandlers.SaveUser)
	route.PUT("/users", routeshandlers.SaveUser)

	route.GET("/prices", routeshandlers.GetAllUsers)
	route.DELETE("/prices/:id", routeshandlers.DeleteUser)
	route.POST("/prices", routeshandlers.SaveUser)
	route.PUT("/prices", routeshandlers.SaveUser)

	route.Run(":9000")
}
