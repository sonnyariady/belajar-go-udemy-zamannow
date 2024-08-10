package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//nilai := "Akang Sonny Kasep"
	nilai2 := "Test satu dua tiga"
	//nilai3 := "QWthbmcgU29ubnkgS2FzZXA="
	encoded := base64.StdEncoding.EncodeToString([]byte(nilai2))
	fmt.Println(encoded)
	//encoded2 := "QWthbmcgU29ubnkgS2FzZXA="
	encoded3 := "QWshbmcgU29ubnkgS2FzZXA=" //salah satu karakter dibedakan dikit
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	decoded2, err2 := base64.StdEncoding.DecodeString(encoded3)
	if err2 != nil {
		fmt.Println("Error:", err2.Error())
	} else {
		fmt.Println(string(decoded2))
	}
}
