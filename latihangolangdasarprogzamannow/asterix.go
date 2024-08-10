package main

import "fmt"

type Alamat struct {
	Jalan, Kota, Negara string
}

func main() {
	alamat1 := Alamat{"Sudirman", "Jaksel", "Indonesia"}
	alamat2 := &alamat1

	alamat2.Kota = "JakTim"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	*alamat2 = Alamat{"SCBD", "Surabaya", "Indonesia"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
