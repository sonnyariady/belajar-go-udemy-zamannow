package main

import "fmt"

func main() {
	avg := RataRata(1, 2, 3, 4, 5)
	fmt.Println("Rata: ", avg)
	angka := []float32{4, 5, 6}
	avg2 := RataRata(angka...)
	fmt.Println("Rata - 2: ", avg2)
}

func RataRata(numbers ...float32) float32 {
	var total float32 = 0
	for _, number := range numbers {
		total = total + number
	}

	return total / float32(len(numbers))
}
