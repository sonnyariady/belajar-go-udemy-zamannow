package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type Alamat struct {
	Nama, Alamat, Telp string
}

func main() {
	var daftaralamats []Alamat
	datacsv := "Nama,Alamat,Telp\n"
	datacsv = datacsv + "Didi,Jl.Mangga,123456\n"
	datacsv = datacsv + "Agus,Jl.Apel,6666\n"
	datacsv = datacsv + "Tono,Jl.Pegangsaan,4444\n"
	fmt.Println(datacsv)

	reader := csv.NewReader(strings.NewReader(datacsv))

	for idx := 0; ; idx++ {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}
		if idx == 0 {
			continue
		}
		/*
			fmt.Println("Index: ", idx)
			fmt.Println(records)
			fmt.Println("Pecah record")
			fmt.Println("Nama : ", records[0])
			fmt.Println("Alamat : ", records[1])
			fmt.Println("Telp : ", records[2])
		*/
		daftaralamats = append(daftaralamats, Alamat{records[0], records[1], records[2]})

		//masukin ke slice of struct

	}

	fmt.Println("Cetak slice daftar alamat")
	fmt.Println(daftaralamats)

	for idx, itm := range daftaralamats {
		fmt.Printf("Data ke-%d --> Nama : %s, Alamat : %s, Telp : %s\n", idx, itm.Alamat, itm.Alamat, itm.Telp)
	}
}
