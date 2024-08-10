package main

import (
	"errors"
	"fmt"
)

func Pembagian(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return a / b, nil
	}
}

func main() {
	hasil, err := Pembagian(27, 0)
	fmt.Println(hasil)

	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Terjadi error: ", err.Error())
	}
}
