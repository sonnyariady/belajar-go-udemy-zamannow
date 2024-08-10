package main

import "fmt"

type Alamat struct {
	Jalan, Kota, Negara string
}

func main() {
	var alamat1 *Alamat = new(Alamat)
	var alamat2 *Alamat = alamat1

	alamat2.Negara = "India"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
