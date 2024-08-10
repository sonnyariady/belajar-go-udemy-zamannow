package main

import "fmt"

func main() {
	var buah [3]string

	buah[0] = "Mangga"
	buah[1] = "Jambu"
	buah[2] = "Nanas"

	fmt.Println(buah[0])
	fmt.Println(buah[1])
	fmt.Println(buah[2])

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)

	var hewans = [4]string{
		"kucing",
		"sapi",
		"kambing",
		"macan",
	}

	fmt.Println(hewans)
	fmt.Println(hewans[1])
	fmt.Println(len(hewans))

	var makanans = [...]string{
		"dimsum",
		"siomay",
		"sate",
		"soto",
		"cilok",
	}

	fmt.Println(len(makanans))
	fmt.Println(makanans)
}
