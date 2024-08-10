package main

import "fmt"

func Selesai() {
	fmt.Println("Sudah semua dijalankan ya")
}

func PesanHalo(nama string) {
	defer Selesai()
	fmt.Println("Hallo ", nama)
	//panic("Aduhaduh")
	//Selesai()
}

func main() {
	PesanHalo("Budi")
}
