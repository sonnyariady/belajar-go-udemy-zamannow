package main

import (
	"fmt"

	sclc "github.com/sonnyariady/go-simple-calc/v2"
)

func main() {
	a := float64(3)
	b := float64(4)
	fmt.Println(sclc.Jumlah(a, b))
	fmt.Println(sclc.Kali(a, b))
	fmt.Println(sclc.Bagi(a, b))
	fmt.Println(sclc.Kurang(a, b))
	fmt.Println(sclc.Pitagoras(a, b))
	fmt.Println(sclc.JumlahKuadrat(a, b, 5))
}
