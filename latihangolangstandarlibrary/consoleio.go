package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	var angka1, angka2 int

	fmt.Print("Input bilangan 1: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	angka1, _ = strconv.Atoi(input)

	fmt.Print("Input bilangan 2: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	angka2, _ = strconv.Atoi(input)

	fmt.Println(angka1, "+", angka2, "=", jumlah(angka1, angka2))
	fmt.Println(angka1, "*", angka2, "=", kali(angka1, angka2))

}

func jumlah(a, b int) int {
	return a + b
}

func kali(a, b int) int {
	return a * b
}
