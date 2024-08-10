package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var now time.Time = time.Now()

	fmt.Println("Now : ", now)
	fmt.Println("Now Local : ", now.Local())
	fmt.Println("Now UTC : ", now.UTC())

	var utc time.Time = time.Date(2022, time.August, 8, 4, 30, 0, 0, time.UTC)
	fmt.Println("UTC :", utc)
	fmt.Println("UTC to Local :", utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2023-03-14 16:11:30"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println("Setelah tambah 40 hari")
	valueTime = valueTime.AddDate(0, 0, 40)
	fmt.Println(valueTime)

	tanggal2 := time.Date(2023, time.May, 21, 5, 45, 0, 0, time.UTC)
	durasi := tanggal2.Sub(valueTime)

	fmt.Println("Durasi penuh: ", durasi)
	fmt.Println("Durasi jam: ", durasi.Hours())
	fmt.Println("Durasi menit: ", durasi.Minutes())
	fmt.Println("Durasi detik: ", durasi.Seconds())

	fmt.Println("Tipe durasi: ", reflect.TypeOf(durasi))
}
