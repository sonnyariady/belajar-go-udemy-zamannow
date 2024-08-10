package main

import "fmt"

type Operasi func(float32, float32) (float32, string)

func main() {

	sayHelloWithFilter("Sonny", spamFilter)
	sayHelloWithFilter("Monyet", spamFilter)

	PergiDenganGerak("Burung", gerakDengan)
	PergiDenganGerak("Ikan", gerakDengan)

	PergiDenganGerak("Burung", gerakMakan)
	PergiDenganGerak("Katak", gerakMakan)

	operasiBilangan(12, 3, Bagi)
	operasiBilangan(12, 3, Kali)
}

func operasiBilangan(a float32, b float32, filter Operasi) {
	hasilop, simbolop := filter(a, b)
	fmt.Println(a, simbolop, b, "=", hasilop)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func PergiDenganGerak(hewan string, filter func(string) string) {
	namaGerak := filter(hewan)
	//fmt.Print(hewan)
	fmt.Println(hewan, "bergerak dengan", namaGerak)
}

func gerakDengan(hewan string) string {
	switch hewan {
	case "Burung":
		return "terbang"
	case "Ikan":
		return "berenang"
	case "Katak":
		return "melompat"
	case "Ular":
		return "merayap"
	default:
		return "jalan biasa"
	}
}

func gerakMakan(hewan string) string {
	switch hewan {
	case "Burung":
		return "cari cacing"
	case "Ikan":
		return "berputar"
	case "Katak":
		return "jilat lidah"
	case "Ular":
		return "melilit mangsa"
	default:
		return "ngesot"
	}
}

func spamFilter(name string) string {
	if name == "Monyet" || name == "Anjing" {
		return "..."
	} else {
		return name
	}

}

func Tambah(a float32, b float32) (float32, string) {
	return a + b, "+"
}

func Kurang(a float32, b float32) (float32, string) {
	return a - b, "-"
}

func Kali(a float32, b float32) (float32, string) {
	return a * b, "*"
}

func Bagi(a float32, b float32) (float32, string) {
	return a / b, "/"
}
