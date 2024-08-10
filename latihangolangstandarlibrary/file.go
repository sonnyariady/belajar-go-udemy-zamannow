package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(nama string, pesan string) error {
	file, err := os.OpenFile(nama, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(pesan)
	return nil
}

func readFile(nama string) (string, error) {
	file, err := os.OpenFile(nama, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var pesan string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		pesan += string(line) + "\n"
	}
	return pesan, nil
}

func appendFile(nama string, pesan string) error {
	file, err := os.OpenFile(nama, os.O_APPEND|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(pesan)
	return nil
}

func main() {
	createNewFile("filesatu.log", "isinya apa saja boleh")

	liriklagu, errbaca := readFile("lirik.lrk")
	if errbaca != nil {
		fmt.Println("Error baca: ", errbaca)
	} else {
		fmt.Println(liriklagu)
	}

	appendFile("filesatu.log", "\ntambah lagi dimari deh")

}
