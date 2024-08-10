package main

import "fmt"

type Alamat struct {
	Jalan, Kota, Negara string
}

func main() {
	alamat1 := Alamat{
		Jalan:  "Mangga",
		Kota:   "Bekasi",
		Negara: "Indonesia",
	}
	alamat2 := alamat1
	alamat2.Kota = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	alamat3 := &alamat1
	alamat3.Kota = "Subang"

	fmt.Println(alamat1)
	fmt.Println(*alamat3)

	var alamat4 *Alamat = &alamat2

	alamat4.Negara = "Jepang"

	fmt.Println(alamat2)
	fmt.Println(alamat4)

	alamat4 = &Alamat{"Jambu", "Tokyo", "Jepang"}

	alamat4.Jalan = "Shiro"

	fmt.Println(alamat2)
	fmt.Println(alamat4)

}
