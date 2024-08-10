package main

import "fmt"

func main() {
	namadepan, namabelakang := NamaPanjang()
	fmt.Println(namadepan + " " + namabelakang)
	luas, keliling := Ukuran(3, 4)
	fmt.Println("Luas: ", luas)
	fmt.Println("Keliling: ", keliling)

	nama1, _ := NamaPanjang()
	fmt.Println(nama1)
}

func NamaPanjang() (string, string) {
	return "Akang", "Sonny"
}

func Ukuran(panjang float32, lebar float32) (float32, float32) {
	return (panjang * lebar), ((2 * panjang) + (2 * lebar))
}
