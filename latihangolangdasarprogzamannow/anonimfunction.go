package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "goblok" || name == "babi"
	}
	registerUser("Sonny", blacklist)
	registerUser("goblok", blacklist)
	registerUser("pecun", func(nama string) bool {
		return nama == "pecun" || nama == "bangke"
	})
	registerUser("pintar", func(nama string) bool {
		return nama == "bego" || nama == "tolol"
	})
}
