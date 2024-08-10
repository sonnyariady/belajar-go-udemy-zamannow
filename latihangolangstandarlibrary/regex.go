package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("Valid Email: ", isValidEmail("sonny@gmail.com"))
	fmt.Println("Valid Email: ", isValidEmail("sonny"))

	var regex = regexp.MustCompile("e([a-z])o")
	var regex2 = regexp.MustCompile("([a-z])ny")
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ado"))
	fmt.Println(regex.FindAllString("ado aku edo eno", 10))
	fmt.Println(regex2.MatchString("donny"))
	fmt.Println(regex2.MatchString("sonny"))
	fmt.Println(regex2.MatchString("robby"))
}

// Fungsi untuk memvalidasi email dengan regex
func isValidEmail(email string) bool {
	// Regex untuk email yang umum
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Kompilasi regex
	re := regexp.MustCompile(emailRegex)

	// Cek apakah email sesuai dengan pola regex
	return re.MatchString(email)
}
