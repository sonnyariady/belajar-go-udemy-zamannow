package main

import (
	"fmt"
	"regexp"
)

// Fungsi untuk memvalidasi password dengan regex
func isValidPassword2(password string) bool {
	// Regex untuk password yang harus memiliki huruf kecil, huruf besar, dan angka
	const passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`

	// Kompilasi regex
	re := regexp.MustCompile(passwordRegex)

	// Cek apakah password sesuai dengan pola regex
	return re.MatchString(password)
}

func main() {
	// Contoh password untuk diuji
	passwords := []string{
		"Password1",
		"password1",
		"PASSWORD1",
		"Pass1",
		"ValidPassword123",
	}

	for _, password := range passwords {
		if isValidPassword2(password) {
			fmt.Printf("%s is a valid password\n", password)
		} else {
			fmt.Printf("%s is an invalid password\n", password)
		}
	}
}
