package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("aku ingin begini\naku ingin begitu\ningin ini ingin itu banyak sekali\n")
	reader := bufio.NewReader(input)

	for idx := 0; ; idx++ {
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println("prefix: ", prefix)
		fmt.Println("Baris ", idx, " = ", string(line))
		//fmt.Printf("baris %d : %s\n", idx, line)
	}
}
