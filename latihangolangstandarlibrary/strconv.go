package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, errBol := strconv.ParseBool("true")

	if errBol == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error: ", errBol.Error())
	}

	flthasil, errflt := strconv.ParseFloat("2.5", 32)

	if errflt == nil {
		fmt.Println(flthasil)
	} else {
		fmt.Println("Error: ", errflt.Error())
	}
	binary := strconv.FormatInt(14, 3)
	fmt.Println(binary)

}
