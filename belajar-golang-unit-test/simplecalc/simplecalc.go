package mysimplecalc

import "math"

func Jumlah(a, b float64) float64 {
	return a + b
}

func Kali(a, b float64) float64 {
	return a * b
}

func Kurang(a, b float64) float64 {
	return a - b
}

func Bagi(a, b float64) float64 {
	return a / b
}

func Pitagoras(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func JumlahKuadrat(a, b, c float64) float64 {
	return math.Pow(a, 2) + math.Pow(b, 2) + math.Pow(c, 2)
}
