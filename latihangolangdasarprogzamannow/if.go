package main

import "fmt"

func main() {
	password := "1234"

	if password == "123" {
		fmt.Println("Benar")
	}

	nama := "Dimas"

	if len(nama) > 5 {
		fmt.Println("Panjang lebih dari 5")
	} else {
		fmt.Println("Panjang kurang atau sama dengan 5")
	}

}
