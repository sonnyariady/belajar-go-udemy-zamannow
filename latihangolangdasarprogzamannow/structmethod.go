package main

import "fmt"

type Mobil struct {
	Nama     string
	Produksi string
	Tahun    int
}

func (mobil Mobil) cetakSpec() {
	fmt.Println("Namanya:", mobil.Nama)
	fmt.Println("Diproduksi oleh:", mobil.Produksi)
	fmt.Println("Dibuat tahun:", mobil.Tahun)
}

func (mobil Mobil) hargaSewa(Jam int) int {
	var hargaSewa int
	var totalHarga int
	switch mobil.Nama {
	case "Xenia":
		hargaSewa = 2000
	case "Kijang":
		hargaSewa = 3000
	default:
		hargaSewa = 1000
	}
	totalHarga = Jam * hargaSewa
	return totalHarga
}

func main() {
	mobil1 := Mobil{"Kijang", "Toyota", 2015}
	mobil2 := Mobil{"Xenia", "Daihatsu", 2013}
	mobil1.cetakSpec()
	mobil2.cetakSpec()
	sewaKijang := mobil1.hargaSewa(4)
	sewaXenia := mobil2.hargaSewa(3)
	fmt.Println("Harga Sewa Kijang:", sewaKijang)
	fmt.Println("Harga Sewa Xenia:", sewaXenia)
}
