package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed appname.txt
var appname string

//go:embed kucing.jpg
var kucingbyte []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(appname)
	err := ioutil.WriteFile("kucing_new.jpg", kucingbyte, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
