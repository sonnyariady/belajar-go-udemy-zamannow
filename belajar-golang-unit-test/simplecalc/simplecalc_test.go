package mysimplecalc

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//before
	//disini bisa ngapain aja seperti melakukan koneksi ke database
	fmt.Println("Sebelum unit test")
	m.Run()
	fmt.Println("Setelah unit test")
}

func TestTableJumlah(t *testing.T) {
	fmt.Println("Mulai table test Jumlah")
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Jumlah3_4",
			request:  "3,4",
			expected: "7",
		},
		{
			name:     "Jumlah19_4",
			request:  "19,4",
			expected: "23",
		},
		{
			name:     "Jumlah7_24",
			request:  "7,24",
			expected: "31",
		},
	}

	for _, test := range tests {
		params := strings.Split(test.request, ",")
		angka1, _ := strconv.ParseFloat(params[0], 64)
		angka2, _ := strconv.ParseFloat(params[1], 64)
		expectedhasil, _ := strconv.ParseFloat(test.expected, 64)
		t.Run(test.name, func(t *testing.T) {
			result := Jumlah(angka1, angka2)
			require.Equal(t, expectedhasil, result)
		})

	}

}

func BenchmarkPitagoras(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pitagoras(float64(3), float64(4))
	}
}

func BenchmarkSubPitagoras(b *testing.B) {
	b.Run("Pitagoras1", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Pitagoras(float64(3), float64(4))
		}
	})

	b.Run("Pitagoras2", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Pitagoras(float64(4), float64(7))
		}
	})
}

func BenchmarkTablePitagoras(b *testing.B) {
	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "Pitagoras3_4",
			request: "3,4",
		},
		{
			name:    "Pitagoras19_4",
			request: "19,4",
		},
		{
			name:    "Pitagoras7_24",
			request: "7,24",
		},
	}

	for _, test := range tests {
		params := strings.Split(test.request, ",")
		angka1, _ := strconv.ParseFloat(params[0], 64)
		angka2, _ := strconv.ParseFloat(params[1], 64)
		b.Run(test.name, func(b *testing.B) {
			Pitagoras(angka1, angka2)
		})

	}
}

func TestSubTest(t *testing.T) {
	t.Run("PitagorasBenar", func(t *testing.T) {
		hasil := Pitagoras(3, 4)
		assert.Equal(t, float64(5), hasil, "Hasilnya harus 5")
		fmt.Println("Assert PitagorasBenar selesai")
	})

	t.Run("PitagorasSalah", func(t *testing.T) {
		hasil := Pitagoras(3, 6)
		assert.Equal(t, 5, hasil, "Hasilnya harus 5")
		fmt.Println("Assert PitagorasSalah selesai")
	})
}

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
	hasil := Pitagoras(3, 4)
	if hasil != 5 {
		t.Error("Hasilnya salah, yg benar ", 5, "tapi hasilnya", hasil)
	}
	fmt.Println("Dieksekusi")
}

func TestPitagorasAssert(t *testing.T) {
	hasil := Pitagoras(3, 4)
	assert.Equal(t, 5, hasil, "Hasilnya harus 5")
	fmt.Println("Assert Pitagoras selesai")
}

func TestPitagorasRequire(t *testing.T) {
	hasil := Pitagoras(3, 4)
	require.Equal(t, 5, hasil, "Hasilnya harus 5")
	fmt.Println("Assert Pitagoras selesai")
}

func TestPitagorasSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa jalan di Windows. Platform Anda adalah ", runtime.GOOS)
	}

	hasil := Pitagoras(3, 4)
	assert.Equal(t, 5, hasil, "Hasilnya harus 5")
	fmt.Println("Assert Pitagoras selesai dengan platform ", runtime.GOOS)
}

func TestPitagoras2(t *testing.T) {
	hasil := Pitagoras(3, 4)
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
