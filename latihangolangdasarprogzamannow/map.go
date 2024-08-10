package main

import "fmt"

func main() {
	ibukota := map[string]string{
		"Indonesia": "Jakarta",
		"Jepang":    "Tokyo",
	}
	ibukota["Rusia"] = "Moskow"
	fmt.Println(ibukota["Jepang"])
	fmt.Println(ibukota["Rusia"])
	fmt.Println(ibukota)
	delete(ibukota, "Jepang")
	fmt.Println(ibukota)

	suarahewan := make(map[string]string)
	suarahewan["kucing"] = "meong"
	suarahewan["harimau"] = "aum"
	fmt.Println(suarahewan)
}
