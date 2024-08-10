package main

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ValidationErrors    = errors.New("Error validasi")
	NotFoundErrors      = errors.New("Error tidak ketemu")
	NotAuthorizedErrors = errors.New("Error tidak terotorisasi")
)

func GetById(id string) error {
	if id == "" {
		return ValidationErrors
	}
	if id == "budi" {
		return NotFoundErrors
	}

	if id == "joko" {
		return NotAuthorizedErrors
	}

	return nil
}

func main() {
	err := GetById("budi")
	if err != nil {
		if errors.Is(err, ValidationErrors) {
			fmt.Println("Error validasi")
		} else if errors.Is(err, NotFoundErrors) {
			fmt.Println("Error not found")
		} else if errors.Is(err, NotAuthorizedErrors) {
			fmt.Println("Error tidak otorisasi")
		} else {
			fmt.Println("Error tak diketahui")
		}
		fmt.Println("Tipe err: ", reflect.TypeOf(err))
	} else {
		fmt.Println("Sukses bro")
	}

}
