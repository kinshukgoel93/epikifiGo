package main

import (
	"epifigo/connection"
	"fmt"
)

func main() {
	fmt.Println("Welcome to EPIFIGO")
	connection.ConnectToMongo()

}
