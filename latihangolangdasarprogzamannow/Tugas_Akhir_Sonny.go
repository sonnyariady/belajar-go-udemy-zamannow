package main

import "fmt"

func main() {
	var makanans = []string{"Ketoprak", "Soto", "Gado-gado", "Sate"}
	cetakArrayMakanan(makanans)
	makanans[1] = "Siomay"
	cetakArrayMakanan(makanans)
}

func cetakArrayMakanan(arr []string) {
	for _, itemarr := range arr {
		switch itemarr {
		case "Ketoprak":
			fmt.Println("Ketupat Tahu Digeprak")
		case "Soto":
			fmt.Println("Pakai jeruk nipis")
		case "Siomay":
			fmt.Println("Jangan pakai pare")
		case "Gado-gado":
			fmt.Println("Pakai kacang panjang")
		default:
			fmt.Println("Minta bungkus")
		}
	}
}
