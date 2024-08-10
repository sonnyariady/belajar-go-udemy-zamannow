package main

import "fmt"

type Alamat struct {
	Jalan, Kota, Negara string
}

func GantiJalan(alamat *Alamat) {
	alamat.Jalan = "Pegangsaan"
}

func main() {
	alamat1 := Alamat{
		Jalan:  "Mangga",
		Kota:   "Bekasi",
		Negara: "Indonesia",
	}

	GantiJalan(&alamat1)

	fmt.Println(alamat1)

}
