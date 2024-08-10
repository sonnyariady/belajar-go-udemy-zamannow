package main

import "fmt"

func main() {
	fmt.Println(Faktorial(2))
	fmt.Println(Faktorial(3))
	fmt.Println(Faktorial(4))
}

func Faktorial(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * Faktorial(angka-1)
	}

}
