package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.4))
	fmt.Println(math.Floor(1.4))
	fmt.Println(math.Floor(1.8))
	fmt.Println(math.Round(1.8))
	fmt.Println(math.Round(1.2))
	fmt.Println(math.Min(19, 11))
	fmt.Println(math.Max(19, 11))
	fmt.Println(math.Sqrt(81))
	fmt.Println(math.Pow(2, 4))
}
