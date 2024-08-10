package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sonny Ariady", "ria"))
	fmt.Println(strings.Contains("Sonny Ariady", "son"))
	fmt.Println(strings.Split("Akang,Sonny,Kasep,Ariady", ","))
	fmt.Println(strings.ToLower("Akang Sonny Kasep"))
	fmt.Println(strings.ToUpper("Akang Sonny Kasep"))
	fmt.Println(strings.ReplaceAll("Akang Sonny Kasep", "a", "i"))
}
