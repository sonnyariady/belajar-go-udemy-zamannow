package main

import "fmt"

type Hewan struct {
	Nama string
}

func (hewan Hewan) GetName() string {
	return hewan.Nama
}

func (hewan Hewan) Berjalan() {
	switch hewan.GetName() {
	case "Kucing":
		fmt.Println("Memanjat")
	case "Ular":
		fmt.Println("Merayap")
	case "Ikan":
		fmt.Println("Berenang")
	case "Katak":
		fmt.Print("Melompat")
	default:
		fmt.Println("Jalan perlahan")
	}
}

type IHewan interface {
	GetName() string
	Berjalan()
}

func SayHello(value IHewan) {
	fmt.Println("Hello", value.GetName())
}

func Gerak(value IHewan) {
	value.Berjalan()
}

func main() {
	hewan1 := Hewan{
		Nama: "Kucing",
	}
	hewan2 := Hewan{"Ular"}

	hewan3 := Hewan{
		Nama: "Ikan",
	}

	SayHello(hewan1)
	Gerak(hewan1)

	SayHello(hewan2)
	Gerak(hewan2)

	SayHello(hewan3)
	Gerak(hewan3)
}
