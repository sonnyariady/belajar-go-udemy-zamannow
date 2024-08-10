package main

import "fmt"

func main() {
	type NPWP string

	var npwpku NPWP = "781234"
	var string1 string = "666"
	var npwp2 NPWP = NPWP(string1)

	fmt.Println(npwpku)
	fmt.Println(npwp2)
	fmt.Println(NPWP("1246"))
}
