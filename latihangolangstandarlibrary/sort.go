package main

import (
	"fmt"
	"sort"
)

type Karwayan struct {
	Nama string
	Umur int
	Gaji float32
}

type KaryawanSlice []Karwayan

func (k KaryawanSlice) Len() int {
	return len(k)
}

/*
	func (k KaryawanSlice) LessUmur(i, j int) bool {
		return k[i].Umur < k[j].Umur
	}
*/
func (k KaryawanSlice) Less(i, j int) bool {
	return k[i].Gaji < k[j].Gaji
}

func (k KaryawanSlice) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func main() {
	list := []Karwayan{
		{"Sonny", 41, 25.4},
		{"Wirya", 22, 28.1},
		{"Doro", 27, 23.8},
	}

	//sort.Sort(KaryawanSlice(list))

	sort.Sort(KaryawanSlice(list))

	fmt.Println(list)
}
