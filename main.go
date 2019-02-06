package main

import (
	"fmt"

	"./config"
	_ "./db"
)

func main() {
	fmt.Printf("%#v\n", config.Config)
}
