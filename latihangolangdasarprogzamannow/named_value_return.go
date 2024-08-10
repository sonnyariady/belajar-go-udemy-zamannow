package main

import "fmt"

func main() {
	firstname, middlename, lastname := NamaLengkap()
	fmt.Println(firstname, middlename, lastname)
	luas, keliling := Ukuran(4, 5)
	fmt.Println("Luas: ", luas)
	fmt.Println("Keliling: ", keliling)
}

func NamaLengkap() (pertama, tengah, akhir string) {
	pertama = "Andre"
	tengah = "Setiawan"
	akhir = "Hutabarat"
	return pertama, tengah, akhir
}

func Ukuran(panjang float32, lebar float32) (luas float32, keliling float32) {
	luas = panjang * lebar
	keliling = (2 * panjang) + (2 * lebar)
	return
}
