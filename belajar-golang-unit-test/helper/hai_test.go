package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkHalloSonny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hallo("Sonny")
	}
}

func TestHalloSonny(t *testing.T) {
	hasil := Hallo("Sonny")
	if hasil != "Hai Sonny" {
		//unit test gagal
		t.Fail()
	}
	fmt.Println("TestHalloSonny done")
}

func TestHalloSonny2(t *testing.T) {
	hasil := Hallo("Sonny")
	if hasil != "Hai Sonny" {
		//unit test gagal
		t.Error("Seharusnya 'Hai Sonny', tetapi hasilnya ", hasil)
	}
	fmt.Println("TestHalloSonny2 done")
}

func TestHalloSonny3(t *testing.T) {
	hasil := Hallo("Sonny")
	if hasil != "Hai Sonny" {
		//unit test gagal
		t.Fatal("Seharusnya 'Hai Sonny', tetapi hasilnya ", hasil)
	}
	fmt.Println("TestHalloSonny3 done")
}

func TestHalloSonnyAssertion(t *testing.T) {
	hasil := Hallo("Sonny")
	assert.Equal(t, "Hai Sonny", hasil)
	fmt.Println("Dieksekusi")
}

func TestHalloAndi(t *testing.T) {
	hasil := Hallo("Andi")
	if hasil != "Hai Andi" {
		//unit test gagal
		t.FailNow()
	}
	fmt.Println("TestHalloAndi done")
}
