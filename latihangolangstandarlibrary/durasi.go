package main

import (
	"fmt"
	"time"
)

func main() {
	var durasi1 time.Duration = 100 * time.Second
	var durasi2 time.Duration = 10 * time.Millisecond
	var durasi3 time.Duration = durasi1 - durasi2
	fmt.Println(durasi3)
	fmt.Println("Durasi detik: ", durasi3.Seconds())
	fmt.Printf("Durasi : %d\n ", durasi3)
}
