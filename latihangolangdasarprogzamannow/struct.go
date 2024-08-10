package main

import "fmt"

type Customer struct {
	Nama   string
	Umur   int
	Alamat string
}

func main() {
	var mycust1 Customer
	var mycust2 Customer
	mycust1.Nama = "Sonny"
	mycust1.Alamat = "Jaticempaka"
	mycust1.Umur = 41

	mycust3 := Customer{"Ahmad", 25, "Bogor"}
	mycust4 := Customer{
		Nama:   "Joko",
		Alamat: "Srengseng",
		Umur:   42,
	}

	fmt.Println(mycust1)
	fmt.Println(mycust2)
	fmt.Println(mycust3)
	fmt.Println(mycust4)

}
