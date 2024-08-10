package mysimplecalc

import (
	"fmt"
	"testing"
)

func TestJumlah(t *testing.T) {
	hasil := Jumlah(4, 2)
	if hasil != 6 {
		t.Error("Hasilnya salah")
	}
}

func TestKali(t *testing.T) {
	hasil := Kali(13, 3)
	if hasil != 39 {
		t.Error("Hasilnya salah")
	}
}

func TestPitagoras1(t *testing.T) {
	hasil := Pitagoras(3, 5)
	if hasil != 5 {
		t.Error("Hasilnya salah, yg benar ", 5, "tapi hasilnya", hasil)
	}
	fmt.Println("Dieksekusi")
}

func TestPitagoras2(t *testing.T) {
	hasil := Pitagoras(3, 5)
	if hasil != 5 {
		t.Fatal("Hasilnya salah, yg benar ", 5, "tapi hasilnya", hasil)
	}
	fmt.Println("Tidak Dieksekusi")
}

func TestJumlahKuadrat(t *testing.T) {
	hasil := JumlahKuadrat(3, 4, 5)
	if hasil != 50 {
		t.Error("Hasilnya salah")
	}
}
