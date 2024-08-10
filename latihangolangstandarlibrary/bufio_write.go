package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Aku ingin begini\n")
	_, _ = writer.WriteString("Aku ingin begitu\n")
	_, _ = writer.WriteString("Aku ingin ini itu banyak sekali\n")
	writer.Flush()
}
