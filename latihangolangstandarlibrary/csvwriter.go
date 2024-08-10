package main

import (
	"encoding/csv"
	"os"
)

type Alamat struct {
	Nama, Alamat, Telp string
}

func main() {
	writer := csv.NewWriter(os.Stdout)
	listalamat := []Alamat{{"Sonny", "Jatiwaringin", "111"}, {"Abdul", "Pondok Bambu", "2222"}, {"Firman", "Rambutan", "3333"}}

	_ = writer.Write([]string{"Nama", "Alamat", "Telp"})
	for _, alamat := range listalamat {
		_ = writer.Write([]string{alamat.Nama, alamat.Alamat, alamat.Telp})
	}
	writer.Flush()

}
