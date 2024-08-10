package main

import "fmt"

func random() interface{} {
	return "Ok"
}

func random2() any {
	return "Apalah"
}

func random3() any {
	return 1
}

func random4() any {
	return 4.5
}

func errHandl() {
	errMsg := recover()
	if errMsg != nil {
		fmt.Println("Terjadi error:", errMsg)
	}
	fmt.Println("Program berakhir")
}

func Cetak(value any) {
	switch result := value.(type) {
	case string:
		fmt.Println("String:", result)
	case int:
		fmt.Println("Int:", result)
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	defer errHandl()
	result1 := random()
	result2 := random2()
	result3 := random3()
	result4 := random4()
	result1String := result1.(string)
	fmt.Println(result1String)

	Cetak(result1)
	Cetak(result2)
	Cetak(result3)
	Cetak(result4)

}
