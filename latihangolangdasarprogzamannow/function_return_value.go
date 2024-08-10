package main

import "fmt"

func main() {
	fmt.Println(Hello("Sonny"))
	fmt.Println(Luas(3, 4))
}

func Hello(nama string) string {
	sapaan := "Hello " + nama
	return sapaan
}

func Luas(panjang float32, lebar float32) float32 {
	return panjang * lebar
}
