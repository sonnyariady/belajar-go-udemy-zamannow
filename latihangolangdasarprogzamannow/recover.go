package main

import "fmt"

func Selesai() {
	errmsg := recover()
	if errmsg != nil {
		fmt.Println("Terjadi error", errmsg)
	}
	fmt.Println("Sudah semua dijalankan ya")
}

func PesanHalo(nama string) {
	defer Selesai()
	fmt.Println("Hallo ", nama)
	if nama == "Joko" {
		panic("Aduhaduh " + nama)
	}
	fmt.Println("Statement aman eksekusi akhir method")
}

func hitungPembagian(a float32, b float32) {
	defer Selesai()
	hasil := a / b
	if b == 0 {
		panic("Pembagian dengan nol")
	} else {
		fmt.Println(a, "/", b, "=", hasil)
	}
}

func main() {
	PesanHalo("Joko")
	PesanHalo("Amin")
	hitungPembagian(12, 3)
	hitungPembagian(8, 0)
}
