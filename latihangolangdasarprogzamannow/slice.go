package main

import "fmt"

func main() {
	// Deklarasi slice tanpa nilai awal

	// Menambahkan elemen ke slice
	s := []int{1, 2, 3}
	s = append(s, 4, 5)

	// Menyalin elemen dari satu slice ke slice lain
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	copy(s2, s1)

	fmt.Println(s)

	hewans := []string{"kucing"}
	hewans = append(hewans, "anjing", "bagong", "ayam")

	fmt.Println(hewans)

	arrtest := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arrtest[5:6])
	fmt.Println(arrtest[2:])
	fmt.Println(arrtest[:6])

	bulans := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	slcbulan := bulans[4:7]

	fmt.Println(slcbulan)
	fmt.Println(cap(slcbulan))

	slcbulan = append(slcbulan, "Ramadan", "Syaban", "Rajab", "RabiulAwal", "RabiulAkhir", "Safar")
	fmt.Println(slcbulan)

	var makanan = [...]string{"Nasi", "Tahu", "Tempe", "Pecel", "Soto"}
	fmt.Println(makanan)

	slicemakan := makanan[2:4]
	fmt.Println(slicemakan)

	slicemakan[0] = "OrekTempe"
	slicemakan[1] = "Boplo"

	fmt.Println(makanan)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	daysslice1 := days[5:7]
	fmt.Println(daysslice1)
	daysslice2 := append(daysslice1, "LiburBaru")
	fmt.Println(daysslice1)
	fmt.Println(daysslice2)
	fmt.Println(days)

	daysslice2[0] = "SabtuLama"
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 3)
	newSlice[0] = "Sonny"
	newSlice[1] = "Ariady"
	fmt.Println(newSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	iniSlice = append(iniSlice, 6, 7)

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
