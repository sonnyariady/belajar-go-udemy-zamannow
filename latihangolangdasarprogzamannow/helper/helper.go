package helper

import "fmt"

var versi = "1.1.0"
var Aplikasi = "Test Golang"

func SelamatDatang(nama string) string {
	return "Selamat Datang," + nama
}

func selamatJalan(nama string) string {
	return "Adios, " + nama
}

func Contoh() {
	fmt.Println(selamatJalan("Budi"))
	fmt.Println(versi)
}
