package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "pwd123", "database password")
	var hostname *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username: ", *username)
	fmt.Println("password: ", *password)
	fmt.Println("hostname: ", *hostname)
	fmt.Println("port: ", *port)
}
