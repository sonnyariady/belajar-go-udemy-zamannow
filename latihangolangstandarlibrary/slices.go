package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Karyawan struct {
	Nama, Jabatan string
	Umur          int
	Gaji          float32
}

func main() {
	names := []string{"Tono", "Andi", "Candra", "Supri", "Farhan"}
	tinggi := []int{160, 150, 147, 165, 180, 140}
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(tinggi))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(tinggi))
	fmt.Println(slices.Contains(names, "Farhan"))
	fmt.Println(slices.Contains(names, "Denny"))

	words1 := []string{"saya", "belajar", "python"}
	fmt.Println("Sebelum reverse: ", words1)
	slices.Reverse(words1)
	fmt.Println("Setelah reverse: ", words1)

	listkaryawan := []Karyawan{}
	listkaryawan = append(listkaryawan, Karyawan{Nama: "Wirya", Umur: 27, Jabatan: ".NET Techlead", Gaji: 27.5})
	listkaryawan = append(listkaryawan, Karyawan{Nama: "Sonny", Umur: 41, Jabatan: ".NET Backend", Gaji: 24.5})
	listkaryawan = append(listkaryawan, Karyawan{Nama: "Emde", Umur: 26, Jabatan: "QA", Gaji: 21})
	listkaryawan = append(listkaryawan, Karyawan{Nama: "Doro", Umur: 25, Jabatan: "Frontend", Gaji: 25.75})
	fmt.Println(listkaryawan)

	slices.SortFunc(listkaryawan, func(a, b Karyawan) int {
		return cmp.Compare(a.Umur, b.Umur)
	})
	fmt.Println("Urut by umur")
	fmt.Println(listkaryawan)

	fmt.Println("Urut by gaji")
	slices.SortFunc(listkaryawan, func(a, b Karyawan) int {
		return cmp.Compare(a.Gaji, b.Gaji)
	})
	fmt.Println(listkaryawan)

	fmt.Println("Urut by nama")
	slices.SortFunc(listkaryawan, func(a, b Karyawan) int {
		return cmp.Compare(a.Nama, b.Nama)
	})
	fmt.Println(listkaryawan)

	fmt.Println("Urut by jabatan")
	slices.SortFunc(listkaryawan, func(a, b Karyawan) int {
		return cmp.Compare(a.Jabatan, b.Jabatan)
	})
	fmt.Println(listkaryawan)

}
