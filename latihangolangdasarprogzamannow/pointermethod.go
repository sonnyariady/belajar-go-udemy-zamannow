package main

import "fmt"

type Pria struct {
	Nama string
}

func (pria *Pria) Married() {
	pria.Nama = "Tuan " + pria.Nama
}

func main() {
	sonny := Pria{"Sonny"}
	sonny.Married()

	fmt.Println(sonny)
}
