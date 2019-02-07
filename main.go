package main

import (
	"./routeshandlers"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", routeshandlers.GetAllUsers)
	route.GET("/users/:id", routeshandlers.DeleteUser)
	route.Run(":9000")
}
