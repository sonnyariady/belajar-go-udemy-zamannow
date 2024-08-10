package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("Iterasi ke-", i)
	}

	hewans := []string{"Kucing", "Anjing", "Sapi", "Kambing"}
	for i := 0; i < len(hewans); i++ {
		fmt.Println("Iterasi ke-", i, hewans[i])
	}

	names := []string{"Akang", "Sonny", "Kasep"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
