package main

import "fmt"

func main() {
	fmt.Println(Jumlah(1, 3))
	objjumlah := Jumlah
	fmt.Println(objjumlah(4, 5))

	mengeong := bersaut
	gonggong := bersaut

	mengeong("meong")
	gonggong("gukguk")

}

func Jumlah(a int, b int) int {
	return a + b
}

func bersaut(kata string) {
	fmt.Println("Hey " + kata)
}
